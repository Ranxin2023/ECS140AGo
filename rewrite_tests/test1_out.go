package main

import (
	"fmt"
	"hw1/expr"
)

func main() {
	var result float64

	result = 2 + expr.ParseAndEval("3", expr.Env{})
	fmt.Printf("%d\n", result)
}
