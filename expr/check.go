// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Adapted for ECS 140A at UC Davis

package expr

import (
	"fmt"
	"strings"
)

//!+Check

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (Literal) Check(vars map[Var]bool) error {
	return nil
}

func (u Unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.Op) {
		return fmt.Errorf("unexpected unary op %q", u.Op)
	}
	return u.X.Check(vars)
}

func (b Binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.Op) {
		return fmt.Errorf("unexpected binary op %q", b.Op)
	}
	if err := b.X.Check(vars); err != nil {
		return err
	}
	return b.Y.Check(vars)
}

//!-Check
