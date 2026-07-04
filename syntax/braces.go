// Copyright (c) 2018, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package syntax

var (
	litLeftBrace  = &Lit{Value: "{"}
	litComma      = &Lit{Value: ","}
	litDots       = &Lit{Value: ".."}
	litRightBrace = &Lit{Value: "}"}
)

// SplitBraces parses brace expansions within a word's literal parts.
// If any valid brace expansions are found, they are replaced with BraceExp nodes,
// and the function returns true.
// Otherwise, the word is left untouched and the function returns false.
//
// For example, a literal word "foo{bar,baz}" will result in a word containing
// the literal "foo", and a brace expansion with the elements "bar" and "baz".
//
// It does not return an error; malformed brace expansions are simply skipped.
// For example, the literal word "a{b" is left unchanged.
func SplitBraces(word *Word) bool { _ = "STUB: not implemented"; return false }

// In the common case where a word has no braces, skip any allocs.

// empty lit

// return {x} to a non-brace

// ParseInt with bit size 64 to ensure consistent behavior on 32-bit platforms.

// increment must be a number

// ParseInt with bit size 64 to ensure consistent behavior on 32-bit platforms.

// are start and end both chars or
// non-chars?

// return broken {x..y[..incr]} to a non-brace

// open braces that were never closed fall back to non-braces
