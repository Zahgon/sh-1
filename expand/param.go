// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package expand

import (
	"mvdan.cc/sh/v3/syntax"
)

func nodeLit(node syntax.Node) string { _ = "STUB: not implemented"; return "" }

// UnsetParameterError is returned when a parameter expansion encounters an
// unset variable and [Config.NoUnset] has been set.
type UnsetParameterError struct {
	Node    *syntax.ParamExp
	Message string
}

func (u UnsetParameterError) Error() string { _ = "STUB: not implemented"; return "" }

func overridingUnset(pe *syntax.ParamExp) bool { _ = "STUB: not implemented"; return false }

func (cfg *Config) paramExp(pe *syntax.ParamExp) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// This is the only parameter expansion that the environment
// interface cannot satisfy.

// true if var has been accessed with * or @ index

// else, elems are already sliced

// nothing to replace

// empty string means '?'; nothing to do there

// Is this even possible? If a user runs into this panic,
// it's most likely a bug we need to fix.

// ${var@a} returns variable attribute flags.
// We use orig (before nameref resolve) for the attributes.

// ${var@A} returns a declare statement that recreates the variable.

// TODO: implement prompt expansion (\u, \h, \w, etc.).

// TODO: implement, like @A but listing keys for assoc arrays.

func removePattern(str, pat string, fromEnd, shortest bool) string {
	_ = "STUB: not implemented"
	return ""
}

// use .* to get the right-most shortest match

// simple suffix

// simple prefix

// no need to check error as Translate returns one

// remove the original pattern (the submatch)

func (cfg *Config) varInd(vr Variable, idx syntax.ArithmExpr) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (cfg *Config) namesByPrefix(prefix string) []string { _ = "STUB: not implemented"; return nil }
