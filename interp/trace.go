package interp

import (
	"bytes"
	"io"

	"mvdan.cc/sh/v3/syntax"
)

// tracer prints expressions like a shell would do if its
// options '-o' is set to either 'xtrace' or its shorthand, '-x'.
type tracer struct {
	buf       bytes.Buffer
	printer   *syntax.Printer
	output    io.Writer
	needsPlus bool
}

func (r *Runner) tracer() *tracer { _ = "STUB: not implemented"; return nil }

// string writes s to tracer.buf if tracer is non-nil,
// prepending "+" if tracer.needsPlus is true.
func (t *tracer) string(s string) { _ = "STUB: not implemented"; return }

func (t *tracer) stringf(f string, a ...any) { _ = "STUB: not implemented"; return }

// expr prints x to tracer.buf if tracer is non-nil,
// prepending "+" if tracer.isFirstPrint is true.
func (t *tracer) expr(x syntax.Node) { _ = "STUB: not implemented"; return }

// flush writes the contents of tracer.buf to the tracer.stdout.
func (t *tracer) flush() { _ = "STUB: not implemented"; return }

// newLineFlush is like flush, but with extra new line before tracer.buf gets flushed.
func (t *tracer) newLineFlush() { _ = "STUB: not implemented"; return }

// reset state

// call prints a command and its arguments with varying formats depending on the cmd type,
// for example, built-in command's arguments are printed enclosed in single quotes,
// otherwise, call defaults to printing with double quotes.
func (t *tracer) call(cmd string, args ...string) { _ = "STUB: not implemented"; return }

// fields may be empty for function () {} declarations

// TODO: only first occurrence of set is not printed, succeeding calls are printed

// should never happen
