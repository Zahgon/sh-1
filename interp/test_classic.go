// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package interp

import (
	"mvdan.cc/sh/v3/syntax"
)

const illegalTok = 0

type testParser struct {
	eof bool
	val string
	rem []string

	err func(err error)
}

func (p *testParser) errf(format string, a ...any) { _ = "STUB: not implemented"; return }

func (p *testParser) next() { _ = "STUB: not implemented"; return }

func (p *testParser) followWord(fval string) *syntax.Word { _ = "STUB: not implemented"; return nil }

func (p *testParser) classicTest(fval string, pastAndOr bool) syntax.TestExpr {
	_ = "STUB: not implemented"
	return *new(syntax.TestExpr)
}

func (p *testParser) testExprBase(fval string) syntax.TestExpr {
	_ = "STUB: not implemented"
	return *new(syntax.TestExpr)
}

// make [ -e ] fall back to [ -n -e ], i.e. use
// the operator as an argument

// testUnaryOp is an exact copy of syntax's.
func testUnaryOp(val string) syntax.UnTestOperator {
	_ = "STUB: not implemented"
	return *new(syntax.UnTestOperator)
}

// testBinaryOp is like syntax's, but with -a and -o, and without =~.
func testBinaryOp(val string) syntax.BinTestOperator {
	_ = "STUB: not implemented"
	return *new(syntax.BinTestOperator)
}
