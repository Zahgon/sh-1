package coreutils

// Error wraps any error returned from the core utilities.
type Error struct {
	err error
}

var (
	_ error                       = &Error{}
	_ interface{ Unwrap() error } = &Error{}
)

func (err *Error) Error() string { _ = "STUB: not implemented"; return "" }

func (err *Error) Unwrap() error { _ = "STUB: not implemented"; return nil }
