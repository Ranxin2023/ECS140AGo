package main

import (
	"fmt"
	"hw1/expr"
)

func ParseAndEval(x string, y expr.Env) float64 {
	return 42.0
}

func main() {
	var result float64

	result = ParseAndEval("1 + 2", expr.Env{})
	fmt.Printf("%d\n", result)
}
