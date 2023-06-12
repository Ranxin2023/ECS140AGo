package main

import (
	"fmt"
	"hw1/expr"
)

func main() {
	var result float64

	x := "1 + 2"
	result = expr.ParseAndEval(x, expr.Env{})
	fmt.Printf("%d\n", result)
}
