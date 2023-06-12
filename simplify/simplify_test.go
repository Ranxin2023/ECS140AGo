package simplify

import (
	"hw1/expr"
	"testing"
)

// !+Simplify
func TestSimplify(t *testing.T) {
	tests := []struct {
		expr string
		env  expr.Env
		want string
	}{
		{"5 + 2", expr.Env{}, "7"},
		{"X", expr.Env{}, "X"},
		{"X", expr.Env{"X": 2}, "2"},
		{"-X", expr.Env{"X": 2}, "-2"},
		{"-X", expr.Env{}, "(-X)"},
		{"--X", expr.Env{"X": 2}, "2"},
		{"Y", expr.Env{"X": 2}, "Y"},
		{"Y*0", expr.Env{"X": 2}, "0"},
		{"Y+0", expr.Env{"X": 2}, "Y"},
		{"Y*1", expr.Env{"X": 2}, "Y"},
		{"0*Y", expr.Env{"X": 2}, "0"},
		{"0+Y", expr.Env{"X": 2}, "Y"},
		{"1*Y", expr.Env{"X": 2}, "Y"},
		{"1*Y*1", expr.Env{"X": 2}, "Y"},
		{"X+2", expr.Env{"X": 2}, "4"},
		{"10 / X", expr.Env{"X": 2}, "5"},
		{"Y+2", expr.Env{"X": 2}, "(Y + 2)"},
		{"3 + 5 + X", expr.Env{"X": 2}, "10"},
		{"X + 3 + 5", expr.Env{}, "((X + 3) + 5)"},
		{"2 + X + 3", expr.Env{}, "((2 + X) + 3)"},
		{"X + 2 + 3", expr.Env{}, "((X + 2) + 3)"},
		{"2 + 3 + X", expr.Env{}, "(5 + X)"},
		{"(X + X) - Y", expr.Env{"X": 2}, "(4 - Y)"},
		{"(X + X) - Y", expr.Env{"Y": 8}, "((X + X) - 8)"},
		{"10 - 1 + X - Y", expr.Env{}, "((9 + X) - Y)"},
		{"X + 3 + 5", expr.Env{}, "((X + 3) + 5)"},
		{"-(X + X) - Y", expr.Env{"X": 2}, "(-4 - Y)"},
		{"(X + X) - Y", expr.Env{"Y": 8}, "((X + X) - 8)"},
		{"10 - 1 + X - Y", expr.Env{}, "((9 + X) - Y)"},
		{"X + 3 + 5", expr.Env{}, "((X + 3) + 5)"},
		{"--(3+5)", expr.Env{"X": 2}, "8"},
		// TODO add more tests for 100% test coverage
	}

	for _, test := range tests {
		e, err := expr.Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		// Run the method
		result := Simplify(e, test.env)

		// Display the result
		got := expr.Format(result)

		// Check the result
		if got != test.want {
			t.Errorf("Simplify(%s) in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}

type Foo int

func (Foo) Eval(env expr.Env) float64 {
	return 0.0
}

func (Foo) Check(vars map[expr.Var]bool) error {
	return nil
}
func TestSimplify_Fail(t *testing.T) {

	func() {
		defer func() {
			if recover() == nil {
				t.Errorf("did not panic, but should\n")
			}
		}()

		var f Foo
		Simplify(f, expr.Env{})
	}()

}

//!-Simplify
