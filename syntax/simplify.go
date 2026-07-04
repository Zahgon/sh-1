// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package syntax

// Simplify modifies a node to remove redundant pieces of syntax, and returns
// whether any changes were made.
//
// The changes currently applied are:
//
//	Remove clearly useless parentheses       $(( (expr) ))
//	Remove dollars from vars in exprs        (($var))
//	Remove duplicate subshells               $( (stmts) )
//	Remove redundant quotes                  [[ "$var" == str ]]
//	Merge negations with unary operators     [[ ! -n $var ]]
//	Use single quotes to shorten literals    "\$foo"
func Simplify(n Node) bool { _ = "STUB: not implemented"; return false }

type simplifier struct {
	modified bool
}

func (s *simplifier) visit(node Node) bool { _ = "STUB: not implemented"; return false }

// Don't inline params, as x[i] and x[$i] mean
// different things when x is an associative
// array; the first means "i", the second "$i".

// don't inline params - same as above.

// unquoting enables globbing

func (s *simplifier) simplifyWord(wps []WordPart) []WordPart { _ = "STUB: not implemented"; return nil }

func (s *simplifier) removeParensArithm(x ArithmExpr) ArithmExpr {
	_ = "STUB: not implemented"
	return *new(ArithmExpr)
}

func (s *simplifier) inlineSimpleParams(x ArithmExpr) ArithmExpr {
	_ = "STUB: not implemented"
	return *new(ArithmExpr)
}

// Not a POSIX-like parameter expansion, or not a valid name without `$`, like $3.

// A complex parameter expansion can't be simplified.
//
// Note that index expressions can't generally be simplified
// either. It's fine to turn ${a[0]} into a[0], but others like
// a[*] are invalid in many shells including Bash.

func (s *simplifier) inlineSubshell(stmts []*Stmt) []*Stmt { _ = "STUB: not implemented"; return nil }

func (s *simplifier) unquoteParams(x TestExpr) TestExpr {
	_ = "STUB: not implemented"
	return *new(TestExpr)
}

func (s *simplifier) removeParensTest(x TestExpr) TestExpr {
	_ = "STUB: not implemented"
	return *new(TestExpr)
}

func (s *simplifier) removeNegateTest(x TestExpr) TestExpr {
	_ = "STUB: not implemented"
	return *new(TestExpr)
}
