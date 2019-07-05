package calculator

import (
	"math"
	"strconv"
	"strings"
)

var operators = map[string]struct {
	prec     int
	rAssoc   bool
	function func(float64, float64) float64
}{
	"^": {4, true, math.Pow},
	"*": {3, false, func(x, y float64) float64 { return x * y }},
	"/": {3, false, func(x, y float64) float64 { return x / y }},
	"+": {2, false, func(x, y float64) float64 { return x + y }},
	"-": {2, false, func(x, y float64) float64 { return x - y }},
}

var functions = map[string]func(float64) float64{
	"sin":  math.Sin,
	"cos":  math.Cos,
	"sqrt": math.Sqrt,
	"tan":  math.Tan,
}

var constants = map[string]float64{
	"e":   math.E,
	"pi":  math.Pi,
	"phi": math.Phi,
}

func isParentheses(token string) bool {
	switch token {
	case "(",
		")":
		return true
	}
	return false
}

func CmdLineInputParsing(input string) []string {
	var output []string
	input = strings.Replace(input, " ", "", -1)
	i := 0
	for j, token := range input {
		token := string(token)
		if _, exists := operators[token]; exists || isParentheses(token) {
			if j == i {
				output = append(output, token)
			} else {
				output = append(output, input[i:j], token)

			}
			i = j + 1
		} else {
			continue
		}
	}
	if input[i:] != "" {
		output = append(output, input[i:])
	}
	return output
}

func ShuntingYardAlgorithm(input []string) []string {
	var stack []string
	var rpn []string
	for _, token := range input {
		switch token {
		case "(":
			stack = append(stack, token)
		case ")":
			for {
				operator := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if operator == "(" {
					break
				}
				rpn = append(rpn, operator)
			}
		default:
			if operator, exists := operators[token]; exists {
				for len(stack) > 0 {
					top := stack[len(stack)-1]
					if prevOp, exists := operators[top]; !exists || prevOp.prec < operator.prec || (prevOp.prec == operator.prec && operator.rAssoc) {
						break
					}
					stack = stack[:len(stack)-1]
					rpn = append(rpn, top)

				}
				stack = append(stack, token)
			} else if _, exists := functions[token]; exists {
				stack = append(stack, token)
			} else {
				rpn = append(rpn, token)
			}
		}
	}
	// drain the stack
	for len(stack) > 0 {
		op := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		rpn = append(rpn, op)
	}
	return rpn
}

func ComputeResult(rpn []string) float64 {
	var result []float64
	for _, token := range rpn {
		if operator, exists := operators[token]; exists {
			// pop y
			y := result[len(result)-1]
			result = result[:len(result)-1]
			// pop x
			x := result[len(result)-1]
			result = result[:len(result)-1]
			x = operator.function(x, y)
			result = append(result, x)
		} else if function, exists := functions[token]; exists {
			x := result[len(result)-1]
			result = result[:len(result)-1]
			x = function(x)
			result = append(result, x)
		} else {
			if value, exists := constants[token]; exists {
				result = append(result, value)
			} else {
				value, _ := strconv.ParseFloat(token, 64)
				result = append(result, value)
			}
		}
	}
	return result[0]
}

func Calculate(input string) float64 {
	tokens := CmdLineInputParsing(input)
	rpn := ShuntingYardAlgorithm(tokens)
	result := ComputeResult(rpn)
	return result
}
