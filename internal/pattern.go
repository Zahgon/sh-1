// Copyright (c) 2026, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package internal

import (
	"mvdan.cc/sh/v3/pattern"
)

// ExtendedPatternMatcher returns a [regexp.Regexp.MatchString]-like function
// to support !(pattern-list) extended patterns where possible.
// It can be used instead of [pattern.Regexp] for narrow use cases.
func ExtendedPatternMatcher(pat string, mode pattern.Mode) (func(string) bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// In the future we could try to support !(pattern) without matching
// the entire input, ensuring we add enough test cases.

// Extended pattern matching operators are always on outside of pathname expansion.

// Handle !(pattern-list) negation: when Regexp returns NegExtglobError,
// match the inner pattern and negate the result.

// extNegatedMatcher handles !(pattern-list) extglob negation.
// Only a single !(...) group with fixed-string prefix and suffix is supported.
func extNegatedMatcher(pat string, groups []pattern.NegExtGlobGroup) (func(string) bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use @(inner) to compile the pattern list, then negate the match.

// prefix and suffix overlap in name
