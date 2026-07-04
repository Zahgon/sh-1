// Copyright (c) 2016, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package syntax

import (
	"io"
	"reflect"
)

// Walk traverses a syntax tree in depth-first order: It starts by calling
// f(node); node must not be nil. If f returns true, Walk invokes f
// recursively for each of the non-nil children of node, followed by
// f(nil).
func Walk(node Node, f func(Node) bool) { _ = "STUB: not implemented"; return }

type nilableNode interface {
	Node
	comparable // pointer nodes, which can be compared to nil
}

func walkNilable[N nilableNode](node N, f func(Node) bool) {
	_ = "STUB: not implemented"
	// nil
	return
}

func walkList[N Node](list []N, f func(Node) bool) { _ = "STUB: not implemented"; return }

func walkComments(list []Comment, f func(Node) bool) {
	_ = "STUB: not implemented"
	// Note that []Comment does not satisfy the generic constraint []Node.
	return
}

// DebugPrint prints the provided syntax tree, spanning multiple lines and with
// indentation. Can be useful to investigate the content of a syntax tree.
func DebugPrint(w io.Writer, node Node) error { _ = "STUB: not implemented"; return nil }

type debugPrinter struct {
	out   io.Writer
	level int
	err   error
}

func (p *debugPrinter) printf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (p *debugPrinter) newline() { _ = "STUB: not implemented"; return }

func (p *debugPrinter) print(x reflect.Value) { _ = "STUB: not implemented"; return }
