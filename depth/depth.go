package depth

import (
	"hw1/expr"
)

// Depth should return the maximum number of AST nodes between the root of the
// tree and any leaf (literal or variable) in the tree.

func Depth(e expr.Expr) uint {
	// TODO: implement this function
	switch e := e.(type) {
	case expr.Var:
		return 1
	case expr.Literal:
		return 1
	case expr.Unary:
		return 1 + Depth(e.X)
	case expr.Binary:
		var left = Depth(e.X)
		var right = Depth(e.Y)
		if left > right {
			return left + 1
		}
		return right + 1
	default:
		panic("illegal type")

	}

}
