// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package interp

import (
	"mvdan.cc/sh/v3/expand"
	"mvdan.cc/sh/v3/syntax"
)

func newOverlayEnviron(parent expand.Environ, background bool) *overlayEnviron {
	_ = "STUB: not implemented"
	return nil
}

// We could do better here if the parent is also an overlayEnviron;
// measure with profiles or benchmarks before we choose to do so.

// overlayEnviron is our main implementation of [expand.WriteEnviron].
type overlayEnviron struct {
	// parent is non-nil if [values] is an overlay over a parent environment
	// which we can safely reuse without data races, such as non-background subshells
	// or function calls.
	parent expand.Environ

	// values maps normalized variable names, per [overlayEnviron.normalize].
	values map[string]namedVariable

	// We need to know if the current scope is a function's scope, because
	// functions can modify global variables. When true, [parent] must not be nil.
	funcScope bool
}

// namedVariable records the original name of a variable for platforms
// where variable names are matched in a case-insensitive way.
type namedVariable struct {
	// TODO(v4): consider adding this field to [expand.Variable],
	// as a general way for a variable to report its original name.
	// This can be useful for GOOS=windows with case insensitive env vars,
	// as otherwise it's not possible to Environ.Get a var
	// and know what was its original name without looping over Environ.Each.
	Name string
	expand.Variable
}

func (o *overlayEnviron) normalize(name string) string { _ = "STUB: not implemented"; return "" }

func (o *overlayEnviron) Get(name string) expand.Variable {
	_ = "STUB: not implemented"
	return *new(expand.Variable)
}

func (o *overlayEnviron) Set(name string, vr expand.Variable) error {
	_ = "STUB: not implemented"
	return nil
}

// Manipulation of a global var inside a function.

// In a function, the parent environment is ours, so it's always read-write.

// unsetting

// modifying the entire variable

func (o *overlayEnviron) Each(f func(name string, vr expand.Variable) bool) {
	_ = "STUB: not implemented"
	return
}

func execEnv(env expand.Environ) []string { _ = "STUB: not implemented"; return nil }

// If a variable is set globally but unset in the
// runner, we need to ensure it's not part of the final
// list. Seems like zeroing the element is enough.
// This is a linear search, but this scenario should be
// rare, and the number of variables shouldn't be large.

func (r *Runner) lookupVar(name string) expand.Variable {
	_ = "STUB: not implemented"
	return *new(expand.Variable)
}

// r.Params may be nil but positional parameters always exist

// not for cryptographic use

// TODO: support setting RANDOM to seed it
// pseudo-random generator from the system

func (r *Runner) envGet(name string) string { _ = "STUB: not implemented"; return "" }

func (r *Runner) delVar(name string) { _ = "STUB: not implemented"; return }

func (r *Runner) setVarString(name, value string) { _ = "STUB: not implemented"; return }

func (r *Runner) setVar(name string, vr expand.Variable) { _ = "STUB: not implemented"; return }

func (r *Runner) setVarWithIndex(prev expand.Variable, name string, index syntax.ArithmExpr, vr expand.Variable) {
	_ = "STUB: not implemented"
	return
}

// When assigning a string to an array, fall back to the
// zero value for the index.

// from the syntax package, we know that value must be a string if index
// is non-nil; nested arrays are forbidden.

// TODO: only clone when inside a subshell and getting a var from outside for the first time

// if the existing variable is already an AssocArray, try our
// best to convert the key to a string

// TODO: only clone when inside a subshell and getting a var from outside for the first time

func (r *Runner) setFunc(name string, body *syntax.Stmt) { _ = "STUB: not implemented"; return }

func stringIndex(index syntax.ArithmExpr) bool { _ = "STUB: not implemented"; return false }

// TODO: make assignVal and [setVar] consistent with the [expand.WriteEnviron] interface

func (r *Runner) assignVal(name string, prev expand.Variable, as *syntax.Assign, valType string) (string, expand.Variable) {
	_ = "STUB: not implemented"
	return "", *new(expand.Variable)
}

// TODO

// don't return the zero value, as that's an unset variable

// Array assignment.

// indexed

// associative

// TODO

// Evaluate values for each array element.

// Index resets our index with a literal value.

// Implicit index, advancing for every word.

// Flatten down the values.

// TODO

// Should only happen if we forgot a case above.
