// Copyright (c) 2021, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package syntax

type QuoteError struct {
	ByteOffset int
	Message    string
}

func (e QuoteError) Error() string { _ = "STUB: not implemented"; return "" }

const (
	quoteErrNull  = "shell strings cannot contain null bytes"
	quoteErrPOSIX = "POSIX shell lacks escape sequences"
	quoteErrRange = "rune out of range"
	quoteErrMksh  = "mksh cannot escape codepoints above 16 bits"
)

// Quote returns a quoted version of the input string,
// so that the quoted version is expanded or interpreted
// as the original string in the given language variant.
//
// Quoting is necessary when using arbitrary literal strings
// as words in a shell script or command.
// Without quoting, one can run into syntax errors,
// as well as the possibility of running unintended code.
//
// An error is returned when a string cannot be quoted for a variant.
// For instance, POSIX lacks escape sequences for non-printable characters,
// and no language variant can represent a string containing null bytes.
// In such cases, the returned error type will be *QuoteError.
//
// The quoting strategy is chosen on a best-effort basis,
// to minimize the amount of extra bytes necessary.
//
// Some strings do not require any quoting and are returned unchanged.
// Those strings can be directly surrounded in single quotes as well.
func Quote(s string, lang LangVariant) (string, error) {
	_ = "STUB: not implemented"

	// Special case; an empty string must always be quoted,
	// as otherwise it expands to zero fields.
	return "", nil
}

// Like regOps; token characters.

// Whitespace; might result in multiple fields.

// Escape sequences would be expanded.

// Would start a comment unless quoted.

// Might result in brace expansion.

// Might result in tilde expansion.

// Might result in globbing.

// Might result in an assignment.

// Nothing to quote; avoid allocating.

// Single quotes are usually best,
// as they don't require any escaping of characters.
// If we have any invalid utf8 or non-printable runes,
// use $'' so that we can escape them.
// Note that we can't use double quotes for those.

// \xXX, fixed at two hexadecimal characters.

// Unfortunately, mksh allows \x to consume more hex characters.
// Ensure that we don't allow it to read more than two.

// Not a valid Unicode code point?

// From the CAVEATS section in R59's man page:
//
// mksh currently uses OPTU-16 internally, which is the same as
// UTF-8 and CESU-8 with 0000..FFFD being valid codepoints.

// \uXXXX, fixed at four hexadecimal characters.

// \UXXXXXXXX, fixed at eight hexadecimal characters.

// Single quotes without any need for escaping.

// The string contains single quotes,
// so fall back to double quotes.

func isHex(r rune) bool { _ = "STUB: not implemented"; return false }
