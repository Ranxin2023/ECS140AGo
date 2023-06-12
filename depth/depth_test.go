package depth

import (
	"hw1/expr"
	"testing"
)

// !+Depth
func TestDepth(t *testing.T) {
	tests := []struct {
		expr     string
		expected uint
	}{
		{"10", 1},
		{"3.14", 1},
		{"-10.0", 2},
		{"xyz", 1},
		{"-10", 2},
		{"-x", 2},
		{"-(3+1)+4", 4},
		{"1-x", 2},
		{"1+x", 2},
		{"1-x", 2},
		{"(1 + 2) + (3 + X)", 3},
		{"(2 * 3) *(4 *5)+ X", 4},
		{"2 + 3 + X", 3},
		{"(1 + 2 + 3 + 4) + X", 5},
		{"1 + 2 + 3 + X", 4},
		{"(3+1)/(3+1)", 3},
		{"3++3", 3},
		// TODO add more tests for 100% test coverage
	}

	for _, test := range tests {
		e, err := expr.Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		// Run the method
		actual := Depth(e)

		// Check the result
		if actual != test.expected {
			t.Errorf("Depth(%s) = %d, expected %d\n", test.expr, actual, test.expected)
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

func TestDepth_Fail(t *testing.T) {

	func() {
		defer func() {
			if recover() == nil {
				t.Errorf("did not panic, but should\n")
			}
		}()

		var f Foo
		Depth(f)
	}()

}

//!-Depth
