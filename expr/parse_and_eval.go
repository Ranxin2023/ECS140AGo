package expr

// Parse and return the eval of the string
func ParseAndEval(s string, env Env) (float64, error) {
	expr, err := Parse(s)
	if err != nil {
		return 0, err
	}

	return expr.Eval(env), nil
}
