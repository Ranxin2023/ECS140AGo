// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Adapted for ECS 140A at UC Davis.

package expr

import (
	"fmt"
	"testing"
)

// !+Eval
func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
		{"-1 + -x", Env{"x": 1}, "-2"},
		{"-1 - x", Env{"x": 1}, "-2"},
		//{"1 + x", Env{"y": 2}, "1+x"},
	}

	for _, test := range tests {
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}

		// Run the method
		result := expr.Eval(test.env)

		// Display the result
		got := fmt.Sprintf("%.6g", result)
		fmt.Printf("\n%s\n", test.expr)
		fmt.Printf("\t%v => %s\n", test.env, got)

		// Check the result
		if got != test.want {
			t.Errorf("\n(%s).Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}

//!-Eval

func TestErrors(t *testing.T) {
	tests := []struct {
		expr    string
		wantErr string
	}{
		{"x % 2", "unexpected '%'"},
		{"math.Pi", "unexpected '.'"},
		{"!true", "unexpected '!'"},
		{`"hello"`, "unexpected '\"'"},
	}

tests:
	for _, test := range tests {
		var err error

		// Try a few different ways to get an error
		for ok := true; ok; ok = false {
			var expr Expr
			expr, err = Parse(test.expr)
			if err != nil {
				//fmt.Println("hi1")
				//fmt.Printf("%v\n", err)
				break
			}

			err = expr.Check(make(map[Var]bool))
			if err != nil {
				//fmt.Println("hi2")
				break
			}

			t.Errorf("unexpected success: %s", test.expr)
			continue tests
		}

		// If we got an error, check if it's the one we expected
		fmt.Printf("Input: %-20s Error Received: %v\n", test.expr, err)
		if err.Error() != test.wantErr {
			t.Errorf("got error %s, want %s", err, test.wantErr)
		}
	}
}
