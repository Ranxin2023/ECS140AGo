// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Adapted for ECS 140A at UC Davis.

// See page 198.

// Package expr provides an expression evaluator.
package expr

import (
	"fmt"
)

//!+env

type Env map[Var]float64

//!-env

//!+Eval1

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l Literal) Eval(_ Env) float64 {
	return float64(l)
}

//!-Eval1

//!+Eval2

func (u Unary) Eval(env Env) float64 {
	switch u.Op {
	case '+':
		return +u.X.Eval(env)
	case '-':
		return -u.X.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.Op))
}

func (b Binary) Eval(env Env) float64 {
	switch b.Op {
	case '+':
		return b.X.Eval(env) + b.Y.Eval(env)
	case '-':
		return b.X.Eval(env) - b.Y.Eval(env)
	case '*':
		return b.X.Eval(env) * b.Y.Eval(env)
	case '/':
		return b.X.Eval(env) / b.Y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.Op))
}

//!-Eval2
