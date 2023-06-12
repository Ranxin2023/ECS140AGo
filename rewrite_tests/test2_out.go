package main

import (
	"fmt"
	"hw1/expr"
)

func main() {
	var result float64

	result = expr.ParseAndEval("1 + 2")
	fmt.Printf("%d\n", result)
}
