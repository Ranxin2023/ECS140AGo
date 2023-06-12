
package main

import (
	"hw1/expr"
	"fmt"
)

func ParseAndEval(x string, y expr.Env) float64 {
	return 42.0
}

func main() {
	var result float64

	// Note the syntax error -- parsing this file should cause a panic().
	result = expr.ParseAndEval("1 + 2", expr.Env{)
	fmt.Printf("%d\n", result)
}
