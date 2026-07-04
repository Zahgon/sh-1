// Copyright (c) 2016, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package syntax

import (
	"unicode/utf8"
)

func asciiLetter[T rune | byte](r T) bool { _ = "STUB: not implemented"; return false }

func asciiDigit[T rune | byte](r T) bool { _ = "STUB: not implemented"; return false }

// bytes that form or start a token
func regOps(r rune) bool { _ = "STUB: not implemented"; return false }

// tokenize these inside parameter expansions
func paramOps(r rune) bool { _ = "STUB: not implemented"; return false }

// tokenize these inside arithmetic expansions
func arithmOps(r rune) bool { _ = "STUB: not implemented"; return false }

func bquoteEscaped(b byte) bool { _ = "STUB: not implemented"; return false }

const escNewl rune = utf8.RuneSelf + 1

func (p *Parser) rune() rune { _ = "STUB: not implemented"; return 0 }

// p.r instead of b so that newline
// character positions don't have col 0.

// Necessary for the last position to be correct.
// TODO: this is not exactly intuitive; figure out a better way.

// Ignore null bytes while parsing, like bash.

// \r\n turns into \n

// \\\r\n turns into \\\n

// TODO: why is this necessary to ensure correct position info?

// We turn backquote command substitutions into $(),
// so we remove the extra backslashes needed by the backquotes.

// we need more bytes to read a full non-ascii rune

// fill reads more bytes from the input src into readBuf.
// Any bytes that had not yet been used at the end of the buffer
// are slid into the beginning of the buffer.
// The number of read bytes is returned, which is at least one
// unless a read error occurred, such as [io.EOF].
func (p *Parser) fill() (n int) { _ = "STUB: not implemented"; return 0 }

// If the reader already gave us [io.EOF], do not try again.
// If we decided to stop for any reason, do not bother reading either.

// don't use p.errPass as we don't want to overwrite p.tok

func (p *Parser) nextKeepSpaces() { _ = "STUB: not implemented"; return }

// Heredocs handle escaped newlines in a special way, but others do not.

func (p *Parser) next() { _ = "STUB: not implemented"; return }

// merge consecutive newline tokens

// If we're parsing $foo#bar, ${foo}#bar, 'foo'#bar, or "foo"#bar,
// #bar is a continuation of the same word, not a comment.
// The same applies inside [[ ]] tests, where '#' has no comment meaning.

// `[` only starts an `[idx]=val` element when it begins a new word;
// otherwise it continues a glob like `foo[0-9]`.

// continuation of open paren

// we are tokenizing manually

// including '(', '|'

// extendedGlob determines whether we're parsing a Bash extended globbing expression.
// For example, whether `*` or `@` are followed by `(` to form `@(foo)`.
func (p *Parser) extendedGlob() bool { _ = "STUB: not implemented"; return false }

// Zsh supports Bash extended globs via the KSH_GLOB option.
// In Bash we would parse extended globs as [ExtGlob] nodes,
// but trying to do that in Zsh would cause ambiguity with glob qualifiers.
// Just like glob qualifiers, parse extended globs as literals in Zsh.

// We don't support e.g. `function @() { ... }` at the moment, but we could.

// NOTE: empty pattern list is a valid globbing syntax like `@()`,
// but we'll operate on the "likelihood" that it is a function;
// only tokenize if its a non-empty pattern list.
// We do this after peeking for just one byte, so that the input `echo *`
// followed by a newline does not hang an interactive shell parser until
// another byte is input.

func (p *Parser) peek() byte { _ = "STUB: not implemented"; return 0 }

func (p *Parser) peekTwo() (byte, byte) {
	_ = "STUB: not implemented"
	// TODO: This should loop for slow readers, e.g. those providing one byte at
	// a time. Use a loop and test it with [testing/iotest.OneByteReader].
	return 0, 0
}

func (p *Parser) regToken(r rune) token { _ = "STUB: not implemented"; return *new(token) }

// Don't call p.rune, as we need to work out p.openBquotes to
// properly handle backslashes in the lexer.

// latter to not tokenise ${$[@]} as $[

// >>&| is an alias for &>>|

// >>&! is an alias for &>>|

// >>& is an alias for &>>

// >&| is an alias for &>|

// >&! is an alias for &>|

func (p *Parser) dqToken(r rune) token { _ = "STUB: not implemented"; return *new(token) }

// Don't call p.rune, as we need to work out p.openBquotes to
// properly handle backslashes in the lexer.

func (p *Parser) paramToken(r rune) token { _ = "STUB: not implemented"; return *new(token) }

// This func gets called by the parser in [runeByRune] mode;
// we need to handle EOF and unexpected runes.

func (p *Parser) arithmToken(r rune) token { _ = "STUB: not implemented"; return *new(token) }

func (p *Parser) newLit(r rune) { _ = "STUB: not implemented"; return }

// don't let r == utf8.RuneSelf go to the second case as [utf8.RuneLen]
// would return -1

func (p *Parser) endLit() (s string) { _ = "STUB: not implemented"; return "" }

func (p *Parser) isLitRedir() bool { _ = "STUB: not implemented"; return false }

func positionalRuneParam[T rune | byte](r T) bool { _ = "STUB: not implemented"; return false }

func singleRuneParam[T rune | byte](r T) bool { _ = "STUB: not implemented"; return false }

func paramNameRune[T rune | byte](r T) bool { _ = "STUB: not implemented"; return false }

func (p *Parser) advanceLitOther(r rune) { _ = "STUB: not implemented"; return }

// escaped byte follows

// zshNumRange peeks at the bytes after '<' to check for a zsh numeric
// range glob pattern like <->, <5->, <-10>, or <5-10>.
func (p *Parser) zshNumRange() bool {
	_ = "STUB: not implemented"
	// Peeking a handful of bytes here should be enough.
	// TODO: This should loop for slow readers, e.g. those providing one byte at
	// a time. Use a loop and test it with [testing/iotest.OneByteReader].
	return false
}

func (p *Parser) advanceLitNone(r rune) { _ = "STUB: not implemented"; return }

// escaped byte follows

// Zsh numeric range glob like <-> or <1-100>; consume until '>'.

func (p *Parser) advanceLitDquote(r rune) { _ = "STUB: not implemented"; return }

// escaped byte follows

func (p *Parser) advanceLitHdoc(r rune) {
	_ = "STUB: not implemented"
	// Unlike the rest of nextKeepSpaces quote states, we handle escaped
	// newlines here. If lastTok==_Lit, then we know we're following an
	// escaped newline, so the first line can't end the heredoc.
	return
}

// escaped byte follows

// This line starts right after an escaped
// newline, so it should never end the heredoc.

// Compare the current line with the stop word.

// minus trailing character

// hit an unexpected EOF or closing backquote

func (p *Parser) quotedHdocWord() *Word { _ = "STUB: not implemented"; return nil }

// Compare the current line with the stop word.

// minus \n

func (p *Parser) advanceLitRe(r rune) { _ = "STUB: not implemented"; return }

func testUnaryOp(val string) UnTestOperator { _ = "STUB: not implemented"; return *new(UnTestOperator) }

func testBinaryOp(val string) BinTestOperator {
	_ = "STUB: not implemented"
	return *new(BinTestOperator)
}
