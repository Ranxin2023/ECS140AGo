package rewrite

import (
	// "fmt"
	"io/ioutil"
	"testing"
)

// !+SimplifyParseAndEval
func TestSimplifyParseAndEval(t *testing.T) {
	tests := []struct {
		in_filename  string
		out_filename string
	}{
		{"test1_in.go", "test1_out.go"},
		{"test2_in.go", "test2_out.go"},
		{"test3_in.go", "test3_out.go"},
		{"test5_in.go", "test5_out.go"},
		{"test6_in.go", "test6_out.go"},
		{"test7_in.go", "test7_out.go"},
		//{"test8_in.go", "test8_out.go"},
	}

	for _, test := range tests {
		in_file_raw, err := ioutil.ReadFile("../rewrite_tests/" + test.in_filename)
		if err != nil {
			panic("Unable to read test input: " + test.in_filename)
		}
		out_file_raw, err := ioutil.ReadFile("../rewrite_tests/" + test.out_filename)
		if err != nil {
			panic("Unable to read test input: " + test.out_filename)
		}

		in_file, out_file := string(in_file_raw), string(out_file_raw)

		rewritten := SimplifyParseAndEval(in_file)

		if out_file != rewritten {
			t.Errorf("SimplifyParseAndEval()[%v] did not match %v\n  %s",
				test.in_filename, test.out_filename, rewritten)
		}
	}
}

func TestSimplifyParseAndEval_Failure(t *testing.T) {
	tests := []struct {
		in_filename string
	}{
		{"test4_in.go"},
	}

	for _, test := range tests {
		in_file_raw, err := ioutil.ReadFile("../rewrite_tests/" + test.in_filename)
		if err != nil {
			panic("Unable to read test input: " + test.in_filename)
		}

		in_file := string(in_file_raw)

		func() {
			defer func() {
				if recover() == nil {
					t.Errorf("SimplifyParseAndEval()[%v] did not panic, but should\n",
						test.in_filename)
				}
			}()

			SimplifyParseAndEval(in_file)
		}()
	}
}

//!-SimplifyParseAndEval
