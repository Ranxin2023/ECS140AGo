package simplify

import (
	"hw1/expr"
)

// Simplify should return the simplified expresion
func Simplify(e expr.Expr, env expr.Env) expr.Expr {
	//TODO implement the simplify
	switch v := e.(type) {
	case expr.Var:
		_, ok := env[v]
		if ok {
			return expr.Literal(v.Eval(env))
		}
		return v
	case expr.Literal:
		return expr.Literal(v.Eval(env))
	case expr.Unary:
		v.X = Simplify(v.X, env)
		switch v.X.(type) {
		case expr.Literal:
			return expr.Literal(v.Eval(env))
		}
		return v
	case expr.Binary:
		v.X = Simplify(v.X, env)
		v.Y = Simplify(v.Y, env)
		switch lv := v.X.(type) {
		case expr.Literal:
			if v.Op == '*' {
				if lv == 0 {
					return expr.Literal(0)
				}
				if lv == 1 {
					return v.Y
				}
			}
			if v.Op == '+' {
				if lv == 0 {
					return v.Y
				}
			}
		}
		switch rv := v.Y.(type) {
		case expr.Literal:
			if v.Op == '*' {
				if rv == 0 {
					return expr.Literal(0)
				}
				if rv == 1 {
					return v.X
				}
			}
			if v.Op == '+' {
				if rv == 0 {
					return v.X
				}
			}
		}
		switch v.X.(type) {
		case expr.Literal:
			switch v.Y.(type) {
			case expr.Literal:
				return expr.Literal(v.Eval(env))
			}
		}
		return v
	default:
		panic("illegal input")
	}

}
