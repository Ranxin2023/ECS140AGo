package rewrite

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"hw1/expr"
	"hw1/simplify"
	"strconv"
	// This package may be helpful...
)

// rewriteCalls should modify the passed AST
func rewriteCalls(node ast.Node) {
	//TODO Write the rewriteCalls function
	ast.Inspect(node, func(node ast.Node) bool {
		switch v := node.(type) {
		case *ast.CallExpr:
			switch fun := v.Fun.(type) {
			case *ast.SelectorExpr:
				if fun.Sel.Name == "ParseAndEval" {
					if len(v.Args) == 2 {
						//first_arg := v.Args[0]
						first_arg_str, ok := v.Args[0].(*ast.BasicLit)
						if !ok || first_arg_str.Kind != token.STRING {
							return true
						}
						first_arg_unquote, _ := strconv.Unquote(first_arg_str.Value)
						/*
							if unquote_err != nil {
								return true
							}
						*/
						expr_arg, expr_err := expr.Parse(first_arg_unquote)
						if expr_err != nil {
							return true
						}
						res := simplify.Simplify(expr_arg, expr.Env{})
						res_str := expr.Format(res)
						first_arg_str.Value = strconv.Quote(res_str)
					}

				}

			}
		}
		return true
	})
}

func SimplifyParseAndEval(src string) string {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	rewriteCalls(f)

	var buf bytes.Buffer
	format.Node(&buf, fset, f)
	return buf.String()
}
