// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

// Package typedjson allows encoding and decoding shell syntax trees as JSON.
// The decoding process needs to know what syntax node types to decode into,
// so the "typed JSON" requires "Type" keys in some syntax tree node objects:
//
//   - The root node
//   - Any node represented as an interface field in the parent Go type
//
// The types of all other nodes can be inferred from context alone.
//
// For the sake of efficiency and simplicity, the "Type" key
// described above must be first in each JSON object.
package typedjson

// TODO: encoding and decoding nodes other than File is untested.

import (
	"io"
	"reflect"

	"mvdan.cc/sh/v3/syntax"
)

// Encode is a shortcut for [EncodeOptions.Encode] with the default options.
func Encode(w io.Writer, node syntax.Node) error { _ = "STUB: not implemented"; return nil }

// EncodeOptions allows configuring how syntax nodes are encoded.
type EncodeOptions struct {
	Indent string // e.g. "\t"

	// Allows us to add options later.
}

// Encode writes node to w in its typed JSON form,
// as described in the package documentation.
func (opts EncodeOptions) Encode(w io.Writer, node syntax.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeValue(val reflect.Value) (reflect.Value, string) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), ""
}

// Construct a new struct with an optional Type, Pos and End,
// and then all the visible fields which aren't positions.

// Node methods are defined on struct pointer receivers.

// posField
// endField

// Do the rest of the fields.

// Addr helps prevent an allocation as we use any fields.

// Encode token-derived operator enums as their syntax string form
// so the wire format stays stable as new tokens are added.

var (
	noValue reflect.Value

	anyType         = reflect.TypeFor[any]()
	anySliceType    = reflect.TypeFor[[]any]()
	posType         = reflect.TypeFor[syntax.Pos]()
	exportedPosType = reflect.TypeFor[*exportedPos]()

	// TODO(v4): derived fields like Type, Pos, and End should have clearly
	// different names to prevent confusion. For example: _type, _pos, _end.
	typeField = reflect.StructField{
		Name: "Type",
		Type: reflect.TypeFor[string](),
		Tag:  `json:",omitempty"`,
	}
	posField = reflect.StructField{
		Name: "Pos",
		Type: exportedPosType,
		Tag:  `json:",omitempty"`,
	}
	endField = reflect.StructField{
		Name: "End",
		Type: exportedPosType,
		Tag:  `json:",omitempty"`,
	}
)

type exportedPos struct {
	Offset, Line, Col uint
}

func encodePos(encPtr reflect.Value, val syntax.Pos) {
	_ = "STUB: not implemented"
	// TODO: perhaps we should encode recovered positions, as that is still useful information.
	return
}

func decodePos(val reflect.Value, enc map[string]any) { _ = "STUB: not implemented"; return }

// Decode is a shortcut for [DecodeOptions.Decode] with the default options.
func Decode(r io.Reader) (syntax.Node, error) {
	_ = "STUB: not implemented"
	return *new(syntax.Node), nil
}

// DecodeOptions allows configuring how syntax nodes are encoded.
type DecodeOptions struct {
	// Empty for now; allows us to add options later.
}

// Decode writes node to w in its typed JSON form,
// as described in the package documentation.
func (opts DecodeOptions) Decode(r io.Reader) (syntax.Node, error) {
	_ = "STUB: not implemented"
	return *new(syntax.Node), nil
}

var nodeByName = map[string]reflect.Type{
	"File": reflect.TypeFor[syntax.File](),
	"Word": reflect.TypeFor[syntax.Word](),

	"Lit":       reflect.TypeFor[syntax.Lit](),
	"SglQuoted": reflect.TypeFor[syntax.SglQuoted](),
	"DblQuoted": reflect.TypeFor[syntax.DblQuoted](),
	"ParamExp":  reflect.TypeFor[syntax.ParamExp](),
	"CmdSubst":  reflect.TypeFor[syntax.CmdSubst](),
	"CallExpr":  reflect.TypeFor[syntax.CallExpr](),
	"ArithmExp": reflect.TypeFor[syntax.ArithmExp](),
	"ProcSubst": reflect.TypeFor[syntax.ProcSubst](),
	"ExtGlob":   reflect.TypeFor[syntax.ExtGlob](),
	"BraceExp":  reflect.TypeFor[syntax.BraceExp](),

	"ArithmCmd":    reflect.TypeFor[syntax.ArithmCmd](),
	"BinaryCmd":    reflect.TypeFor[syntax.BinaryCmd](),
	"IfClause":     reflect.TypeFor[syntax.IfClause](),
	"ForClause":    reflect.TypeFor[syntax.ForClause](),
	"WhileClause":  reflect.TypeFor[syntax.WhileClause](),
	"CaseClause":   reflect.TypeFor[syntax.CaseClause](),
	"Block":        reflect.TypeFor[syntax.Block](),
	"Subshell":     reflect.TypeFor[syntax.Subshell](),
	"FuncDecl":     reflect.TypeFor[syntax.FuncDecl](),
	"TestClause":   reflect.TypeFor[syntax.TestClause](),
	"DeclClause":   reflect.TypeFor[syntax.DeclClause](),
	"LetClause":    reflect.TypeFor[syntax.LetClause](),
	"TimeClause":   reflect.TypeFor[syntax.TimeClause](),
	"CoprocClause": reflect.TypeFor[syntax.CoprocClause](),
	"TestDecl":     reflect.TypeFor[syntax.TestDecl](),

	"UnaryArithm":  reflect.TypeFor[syntax.UnaryArithm](),
	"BinaryArithm": reflect.TypeFor[syntax.BinaryArithm](),
	"ParenArithm":  reflect.TypeFor[syntax.ParenArithm](),

	"UnaryTest":  reflect.TypeFor[syntax.UnaryTest](),
	"BinaryTest": reflect.TypeFor[syntax.BinaryTest](),
	"ParenTest":  reflect.TypeFor[syntax.ParenTest](),

	"WordIter":   reflect.TypeFor[syntax.WordIter](),
	"CStyleLoop": reflect.TypeFor[syntax.CStyleLoop](),
}

func decodeValue(val reflect.Value, enc any) error { _ = "STUB: not implemented"; return nil }

// Type is already used above. Pos and End came from method calls.

// TODO: don't panic on bad input

// Note that encoding/json defaults to float64 for numbers.
