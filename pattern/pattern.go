// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

// Package pattern allows working with shell pattern matching notation, also
// known as wildcards or globbing.
//
// For reference, see
// https://pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap02.html#tag_18_13.
package pattern

import (
	"strings"
)

// Mode can be used to supply a number of options to the package's functions.
// Not all functions change their behavior with all of the options below.
type Mode uint

type SyntaxError struct {
	msg string
	err error
}

func (e SyntaxError) Error() string { _ = "STUB: not implemented"; return "" }

func (e SyntaxError) Unwrap() error {
	_ = "STUB: not implemented"

	// NegExtGlobGroup represents the byte offset range of a single !(expr) group
	// within a pattern string. Start is the offset of '!', End is one past ')'.
	return nil
}

type NegExtGlobGroup struct {
	Start, End int
}

// NegExtGlobError is returned by [Regexp] when an extglob negation operator
// !(pattern-list) is encountered, as Go's [regexp] package does not support
// negative lookahead. Callers can handle this by negating the result of
// matching the inner pattern.
type NegExtGlobError struct {
	Groups []NegExtGlobGroup
}

func (e *NegExtGlobError) Error() string { _ = "STUB: not implemented"; return "" }

// TODO(v4): flip NoGlobStar to be opt-in via GlobStar, matching bash
// TODO(v4): flip EntireString to be opt-out via PartialMatch, as EntireString causes subtle bugs when forgotten
// TODO(v4): rename NoGlobCase to CaseInsensitive for readability

const (
	Shortest          Mode = 1 << iota // prefer the shortest match.
	Filenames                          // "*" and "?" don't match slashes; only "**" does; only makes sense with EntireString too
	EntireString                       // match the entire string using ^$ delimiters
	NoGlobCase                         // do case-insensitive match (that is, use (?i) in the regexp); shopt "nocaseglob"
	NoGlobStar                         // do not support "**"; negated shopt "globstar"
	GlobLeadingDot                     // let wildcards match leading dots in filenames; shopt "dotglob"
	ExtendedOperators                  // support extended pattern matching operators; shopt "extglob" for pathname expansion
)

// Regexp turns a shell pattern into a regular expression that can be used with
// [regexp.Compile]. It will return an error if the input pattern was incorrect.
// Otherwise, the returned expression can be passed to [regexp.MustCompile].
//
// For example, Regexp(`foo*bar?`, true) returns `foo.*bar.`.
//
// Note that this function (and [QuoteMeta]) should not be directly used with file
// paths if Windows is supported, as the path separator on that platform is the
// same character as the escaping character for shell patterns.
func Regexp(pat string, mode Mode) (string, error) {
	_ = "STUB: not implemented"
	// If there are no special pattern matching or regular expression characters,
	// and we don't need to insert extras for the modes affecting non-special characters,
	// we can directly return the input string as a short-cut.
	return "", nil
}

// including those that need escaping since they are
// regular expression metacharacters

// Enable matching `\n` with the `.` metacharacter as globs match `\n`

// stringLexer helps us tokenize a pattern string.
// Note that we can use the null byte '\x00' to signal "no character" as shell strings cannot contain null bytes.
type stringLexer struct {
	s string
	i int
}

func (sl *stringLexer) next() rune { _ = "STUB: not implemented"; return 0 }

func (sl *stringLexer) last() rune { _ = "STUB: not implemented"; return 0 }

func (sl *stringLexer) peekNext() rune { _ = "STUB: not implemented"; return 0 }

func (sl *stringLexer) peekRest() string { _ = "STUB: not implemented"; return "" }

func regexpNext(sb *strings.Builder, sl *stringLexer, mode Mode) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle extended pattern matching operators separately,
// given that they can be one of many two-character prefixes.
// Note that we recurse into the same function in a loop,
// as each of the patterns in the list separated by '|' is a regular pattern.

// position of the operator
// (

// extended operators support a list of "or" separated expressions

// )

// @( is [syntax.GlobOne] for matching once; no suffix needed

// * - matches anything when not in filename mode

// "**" only acts as globstar if it is alone as a path element.

// ** - match any number of slashes or "*" path elements

// **/ - like "**" but requiring a trailing slash when matching

// wrap the expression to ensure that any match has a slash suffix

// with GlobLeadingDot (dotglob), match anything at all

// foo**, **bar, or NoGlobStar - behaves like "*" below

// * - matches anything except slashes and leading dots

// with GlobLeadingDot (dotglob), match anything except slashes

// TODO: surely char classes can be mixed with others, e.g. [[:foo:]xyz]

// TODO: what about overlapping ranges, like: [a--z]

func charClass(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// HasMeta returns whether a string contains any unescaped pattern
// metacharacters: '*', '?', or '['. When the function returns false, the given
// pattern can only match at most one string.
//
// For example, HasMeta(`foo\*bar`) returns false, but HasMeta(`foo*bar`)
// returns true.
//
// This can be useful to avoid extra work, like [Regexp]. Note that this
// function cannot be used to avoid [QuoteMeta], as backslashes are quoted by
// that function but ignored here.
//
// The [Mode] parameter is unused, and will be removed in v4.
func HasMeta(pat string, mode Mode) bool { _ = "STUB: not implemented"; return false }

// QuoteMeta returns a string that quotes all pattern metacharacters in the
// given text. The returned string is a pattern that matches the literal text.
//
// For example, QuoteMeta(`foo*bar?`) returns `foo\*bar\?`.
//
// The [Mode] parameter is unused, and will be removed in v4.
func QuoteMeta(pat string, mode Mode) string { _ = "STUB: not implemented"; return "" }

// short-cut without a string copy
