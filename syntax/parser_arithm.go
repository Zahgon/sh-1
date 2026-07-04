package syntax

// compact specifies whether we allow spaces between expressions.
// This is true for let
func (p *Parser) arithmExpr(compact bool) ArithmExpr {
	_ = "STUB: not implemented"
	return *new(ArithmExpr)
}

// These function names are inspired by Bash's expr.c

func (p *Parser) arithmExprComma(compact bool) ArithmExpr {
	_ = "STUB: not implemented"
	return *new(ArithmExpr)
}

func (p *Parser) arithmExprAssign(compact bool) ArithmExpr {
	_ = "STUB: not implemented"
	// Assign is different from the other binary operators because it's
	// right-associative and needs to check that it's placed after a name
	return *new(ArithmExpr)
}

func (p *Parser) arithmExprTernary(compact bool) ArithmExpr {
	_ = "STUB: not implemented"
	return *new(ArithmExpr)
}

func (p *Parser) arithmExprLor(compact bool) ArithmExpr {
	_ = "STUB: not implemented"
	return *new(ArithmExpr)
}

func (p *Parser) arithmExprLand(compact bool) ArithmExpr {
	_ = "STUB: not implemented"
	return *new(ArithmExpr)
}

func (p *Parser) arithmExprBor(compact bool) ArithmExpr {
	_ = "STUB: not implemented"
	return *new(ArithmExpr)
}

func (p *Parser) arithmExprBxor(compact bool) ArithmExpr {
	_ = "STUB: not implemented"
	return *new(ArithmExpr)
}

func (p *Parser) arithmExprBand(compact bool) ArithmExpr {
	_ = "STUB: not implemented"
	return *new(ArithmExpr)
}

func (p *Parser) arithmExprEquality(compact bool) ArithmExpr {
	_ = "STUB: not implemented"
	return *new(ArithmExpr)
}

func (p *Parser) arithmExprComparison(compact bool) ArithmExpr {
	_ = "STUB: not implemented"
	return *new(ArithmExpr)
}

func (p *Parser) arithmExprShift(compact bool) ArithmExpr {
	_ = "STUB: not implemented"
	return *new(ArithmExpr)
}

func (p *Parser) arithmExprAddition(compact bool) ArithmExpr {
	_ = "STUB: not implemented"
	return *new(ArithmExpr)
}

func (p *Parser) arithmExprMultiplication(compact bool) ArithmExpr {
	_ = "STUB: not implemented"
	return *new(ArithmExpr)
}

func (p *Parser) arithmExprPower(compact bool) ArithmExpr {
	_ = "STUB: not implemented"
	// Power is different from the other binary operators because it's right-associative
	return *new(ArithmExpr)
}

func (p *Parser) arithmExprUnary(compact bool) ArithmExpr {
	_ = "STUB: not implemented"
	return *new(ArithmExpr)
}

func (p *Parser) arithmExprValue(compact bool) ArithmExpr {
	_ = "STUB: not implemented"
	return *new(ArithmExpr)
}

// we want real nil, not (*Word)(nil) as that
// sets the type to non-nil and then x != nil

// nextArith consumes a token.
// It returns true if compact and the token was followed by spaces
func (p *Parser) nextArith(compact bool) bool { _ = "STUB: not implemented"; return false }

func (p *Parser) nextArithOp(compact bool) { _ = "STUB: not implemented"; return }

// arithmExprBinary is used for all left-associative binary operators
func (p *Parser) arithmExprBinary(compact bool, nextOp func(bool) ArithmExpr, operators ...BinAritOperator) ArithmExpr {
	_ = "STUB: not implemented"
	return *new(ArithmExpr)
}

func isArithName(left ArithmExpr) bool { _ = "STUB: not implemented"; return false }

func (p *Parser) followArithm(ftok token, fpos Pos) ArithmExpr {
	_ = "STUB: not implemented"
	return *new(ArithmExpr)
}

func (p *Parser) peekArithmEnd() bool { _ = "STUB: not implemented"; return false }

func (p *Parser) arithmMatchingErr(pos Pos, left, right token) { _ = "STUB: not implemented"; return }

func (p *Parser) matchedArithm(lpos Pos, left, right token) { _ = "STUB: not implemented"; return }

func (p *Parser) arithmEnd(ltok token, lpos Pos, old saveState) Pos {
	_ = "STUB: not implemented"
	return *new(Pos)
}
