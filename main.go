package main

import (
	"github.com/josibake/shuntingyard"
	"html/template"
	"log"
	"net/http"
)

type Calculation struct {
	Input  string
	Result float64
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	result := shuntingyard.Calculate(r.Form["infix"][0])
	calculation := Calculation{r.Form["infix"][0], result}

	t, _ := template.ParseFiles("results.html")
	t.ExecuteTemplate(w, "results.html", calculation)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/calculate", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
