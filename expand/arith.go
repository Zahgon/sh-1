// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package expand

import (
	"mvdan.cc/sh/v3/syntax"
)

// TODO(v4): the arithmetic APIs should return int64 for portability with 32-bit systems,
// even if Bash only supports native int sizes.

func Arithm(cfg *Config, expr syntax.ArithmExpr) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// recursively fetch vars

// default to 0

// TernColon can't happen here

// must have Op==TernColon

func oneIf(b bool) int { _ = "STUB: not implemented"; return 0 }

// atoi is like [strconv.ParseInt](s, BASE, 64), but it handles integer
// base prefixes according to bash-shell's rules, ignores errors, and
// trims whitespace.
//
// For more information about bash's integer base handling syntax,
// refer to the bash manual:
// https://www.man7.org/linux/man-pages/man1/bash.1.html
func atoi(s string) int64 { _ = "STUB: not implemented"; return 0 }

func (cfg *Config) assgnArit(b *syntax.BinaryArithm) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func intPow(a, b int) int { _ = "STUB: not implemented"; return 0 }

func binArit(op syntax.BinAritOperator, x, y int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// x is executed but its result discarded
