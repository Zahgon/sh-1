// Copyright (c) 2017, Andrey Nering <andrey.nering@gmail.com>
// See LICENSE for licensing information

//go:build !unix

package interp

import (
	"context"

	"mvdan.cc/sh/v3/syntax"
)

func mkfifo(path string, mode uint32) error { _ = "STUB: not implemented"; return nil }

// access attempts to emulate [unix.Access] on Windows.
// Windows seems to have a different system of permissions than Unix,
// so for now just rely on what [io/fs.FileInfo] gives us.
func (r *Runner) access(ctx context.Context, path string, mode uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// unTestOwnOrGrp panics. Under Unix, it implements the -O and -G unary tests,
// but under Windows, it's unclear how to implement those tests, since Windows
// doesn't have the concept of a file owner, just ACLs, and it's unclear how
// to map the one to the other.
func (r *Runner) unTestOwnOrGrp(ctx context.Context, op syntax.UnTestOperator, x string) bool {
	_ = "STUB: not implemented"
	return false
}

// waitStatus is a no-op on plan9 and windows.
type waitStatus struct{}

func (waitStatus) Signaled() bool { _ = "STUB: not implemented"; return false }
func (waitStatus) Signal() int    { _ = "STUB: not implemented"; return 0 }
