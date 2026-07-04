// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package interp

import (
	"context"
	"os"

	"mvdan.cc/sh/v3/syntax"
)

// non-empty string is true, empty string is false
func (r *Runner) bashTest(ctx context.Context, expr syntax.TestExpr, classic bool) string {
	_ = "STUB: not implemented"
	return ""
}

// In the classic "test" mode, we already expanded and
// split the list of words, so don't redo that work.

// test, [

// [[

func (r *Runner) binTest(ctx context.Context, op syntax.BinTestOperator, x, y string) bool {
	_ = "STUB: not implemented"
	return false
}

// -ot is the mirror of -nt, so swap the operands and share the logic.

// True if the first operand exists and the second does not,
// or if both exist and the first is newer.

// Should only happen if we forgot a case above.

func (r *Runner) statMode(ctx context.Context, name string, mode os.FileMode) bool {
	_ = "STUB: not implemented"
	return false
}

// These are copied from x/sys/unix as we can't import it here.
const (
	access_R_OK = 0x4
	access_W_OK = 0x2
	access_X_OK = 0x1
)

func (r *Runner) unTest(ctx context.Context, op syntax.UnTestOperator, x string) bool {
	_ = "STUB: not implemented"
	return false
}

// Support [os.File.Fd] methods such as the one on [*os.File].

// TODO: allow term.IsTerminal here too if running in the
// "single process" mode.

// Should only happen if we forgot a case above.
