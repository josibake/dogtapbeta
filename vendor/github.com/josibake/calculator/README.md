[![Coverage Status](https://coveralls.io/repos/github/josibake/learn-tdd/badge.svg?branch=master)](https://coveralls.io/github/josibake/learn-tdd?branch=master)
[![Build Status](https://travis-ci.com/josibake/learn-tdd.svg?branch=master)](https://travis-ci.com/josibake/learn-tdd)

lil command line math utility - purpose of this project is to..

* learn test driven development principles (TDD)
* write a command line utility that has 100% coverage 
* set up a CD/CI pipeline that uses testing before adding new featues

## interesting finds along the way

turns out, parsing an expression is not a trivial problem. the first thing i stumbled across was the shunting yard algorithm, invented by dijkstra. this will be the first implementation. also found this article interesting (adding it here to read it later: http://www.engr.mun.ca/~theo/Misc/exp_parsing.htm). it proposes parsing via recursive descent and he derives a technique known as precedence climbing 

## shunting yard algorithm
