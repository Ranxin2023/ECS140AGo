// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Adapted for ECS 140A at UC Davis.

package expr

// An Expr is an arithmetic expression.
type Expr interface {
	// Eval returns the value of this Expr under the environment env.
	Eval(env Env) float64

	// Check reports errors in this Expr and adds its Vars to the set.
	Check(vars map[Var]bool) error
}

//!+ast

// A Var identifies a variable, e.g., x.
type Var string

// Literal is a numeric constant, e.g., 3.141.
type Literal float64

// Unary represents a unary operator expression, e.g., -x.
type Unary struct {
	Op rune // one of '+', '-'
	X  Expr
}

// Binary represents a binary operator expression, e.g., x+y.
type Binary struct {
	Op   rune // one of '+', '-', '*', '/'
	X, Y Expr
}

//!-ast
