// Copyright (c) 2018, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package expand

import (
	"iter"

	"mvdan.cc/sh/v3/syntax"
)

// Braces performs brace expansion on a word, given that it contains any
// [syntax.BraceExp] parts. For example, the word with a brace expansion
// "foo{bar,baz}" will return two literal words, "foobar" and "foobaz".
//
// Note that the resulting words may share word parts.
//
// Deprecated: use [BracesSeq], which yields words lazily and reports an
// error rather than letting a large sequence allocate huge amounts.
func Braces(word *syntax.Word) []*syntax.Word { _ = "STUB: not implemented"; return nil }

// BracesSeq performs brace expansion on a word, given that it contains any
// [syntax.BraceExp] parts. For example, the word with a brace expansion
// "foo{bar,baz}" will return two literal words, "foobar" and "foobaz".
//
// The iteration yields an error and stops if the total expansion is too
// large, including combinatorial blow-ups across multiple brace expansions
// like {1..100}{1..100}{1..100}. This may be configurable with cfg in the
// future; the parameter is entirely unused for now.
//
// Note that the resulting words may share word parts.
func BracesSeq(cfg *Config, word *syntax.Word) iter.Seq2[*syntax.Word, error] {
	_ = "STUB: not implemented"
	return nil
}

// 16Ki expanded elements is more than any script should need in practice,
// but it's small enough where we don't waste too much memory and CPU.

// bracesSeqRec yields each fully-expanded word descended from word.
// It returns false if iteration should stop.
func bracesSeqRec(word *syntax.Word, yield func(*syntax.Word) bool) bool {
	_ = "STUB: not implemented"
	return false
}

// Yield each word produced by recursing on `next`,
// after prepending `left` to its Parts.

// ParseInt with bit size 64 to ensure consistent behavior on 32-bit platforms.

// ParseInt with bit size 64 to ensure consistent behavior on 32-bit platforms.

func extraLeadingZeros(s string) int { _ = "STUB: not implemented"; return 0 }

// "0" has no extra leading zeros
