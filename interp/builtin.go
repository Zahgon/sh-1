// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package interp

import (
	"bufio"
	"context"

	"mvdan.cc/sh/v3/syntax"
)

// TODO: given the categories below, perhaps this should be more like:
//
//   func IsBuiltin(lang syntax.LangVariant, name string) bool
//
// or perhaps some API that also lets the user iterate through the builtins?
//
// Also, should we move this to the syntax package too?
// It's not a syntactical property strictly speaking,
// but it's also odd to require importing the interp package for it.

// IsBuiltin returns true if the given word is a POSIX Shell
// or Bash builtin.
func IsBuiltin(name string) bool {
	_ = "STUB: not implemented"

	// POSIX Shell builtins, from section 1.d obtained in September 2025 from:
	// https://pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap02.html#tag_18_09_01_01
	return false
}

// POSIX Shell special built-ins, obtained in September 2025 from:
// https://pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap02.html#tag_18_14

// NOTE: our parser treats this as a keyword
// NOTE: our parser treats this as a keyword

// Bash built-ins which are not present in POSIX, obtained in September 2025 from:
// https://man.archlinux.org/man/bash.1.en#SHELL_BUILTIN_COMMANDS

// NOTE: our parser treats this as a keyword
// NOTE: our parser treats this as a keyword

// TODO: surely this is POSIX? but why is it not in the main POSIX spec page?

// NOTE: our parser treats this as a keyword

// TODO: surely this is POSIX? but why is it not in the main POSIX spec page?

// NOTE: an alias for "test", not explicitly listed

// TODO: atoi is duplicated in the expand package.

// atoi is like [strconv.ParseInt](s, 10, 64), but it ignores errors and trims whitespace.
func atoi(s string) int64 { _ = "STUB: not implemented"; return 0 }

type errBuiltinExitStatus exitStatus

func (e errBuiltinExitStatus) Error() string { _ = "STUB: not implemented"; return "" }

// Builtin allows [ExecHandlerFunc] implementations to execute any builtin,
// which can be useful for an exec handler to wrap or combine builtin calls.
//
// Note that a non-nil error may be returned in cases where the builtin
// alters the control flow of the runner, even if the builtin did not fail.
// For example, this is the case with `exit 0` or `return`.
func (hc HandlerContext) Builtin(ctx context.Context, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Runner) builtin(ctx context.Context, pos syntax.Pos, name string, args []string) (exit exitStatus) {
	_ = "STUB: not implemented"
	return *new(exitStatus)
}

// default

// perhaps overly dramatic?

// replicate the commonly implemented behavior of `cd -`
// ref: https://www.man7.org/linux/man-pages/man1/cd.1p.html#OPERANDS

// Note that "wait" without arguments always returns exit status zero.

// TODO: implement. for now, having this as a no-op is better than nothing.

// If the script was not found in PATH or there was any error, pass
// the source path to the open handler so it has a chance to look
// at files it manages (eg: virtual filesystem), and also allow
// it to look for the sourced script in the current directory.

// Keep the current versions of some fields we might modify.

// If we run "source file args...", set said args as parameters.
// Otherwise, keep the current parameters.

// We want to track if the sourced file explicitly sets the
// parameters.

// know that we're inside a sourced script.

// If we modified the parameters and the sourced file didn't
// explicitly set them, we restore the old ones.

// TODO: Consider unix.Exec, i.e. actually replacing
// the process. It's in theory what a shell should do,
// but in practice it would kill the entire Go process
// and it's not available on Windows.

// Note that on Windows, syscall.Stdin is of type uintptr.

// read -a arrayname: split line into fields and assign to indexed array.

// Use -1 as max to get all fields without joining the last ones.

// We can get data back from readLine and an error at the same time, so
// check err after we process the data.

// ""

// TODO: parse any CallExpr perhaps, or even any Stmt

// default signal

// Print non-default signals

// assume it's a signal, the default will be restored

// For now, treat both empty and - the same since ERR and EXIT have no
// default callback.

// Remove the delim from each line read

// Bash sets the delim to an ASCII NUL if provided with an empty
// string.

// mapfileSplit returns a suitable Split function for a [bufio.Scanner];
// the code is mostly stolen from [bufio.ScanLines].
func mapfileSplit(delim byte, dropDelim bool) bufio.SplitFunc {
	_ = "STUB: not implemented"
	return *new(bufio.SplitFunc)
}

// We have a full newline-terminated line.

// If we're at EOF, we have a final, non-terminated line. Return it.

// Request more data.

func (r *Runner) printOptLine(name string, enabled, supported bool) {
	_ = "STUB: not implemented"
	return
}

func (r *Runner) readLine(ctx context.Context, raw bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The AfterFunc was started.
// Wait for it to complete, and reset the file's deadline.

// line continuation

func (r *Runner) changeDir(ctx context.Context, cmd, path string) uint8 {
	_ = "STUB: not implemented"
	return 0
}

func absPath(dir, path string) string { _ = "STUB: not implemented"; return "" }

// TODO: this clean is likely unnecessary

func (r *Runner) absPath(path string) string { _ = "STUB: not implemented"; return "" }

// flagParser is used to parse builtin flags.
//
// It's similar to the getopts implementation, but with some key differences.
// First, the API is designed for Go loops, making it easier to use directly.
// Second, it doesn't require the awkward ":ab" syntax that getopts uses.
// Third, it supports "-a" flags as well as "+a".
type flagParser struct {
	current   string
	remaining []string
}

func (p *flagParser) more() bool { _ = "STUB: not implemented"; return false }

// We're still parsing part of "-ab".

// Nothing left.

// We explicitly stop parsing flags.

// The next argument is not a flag.

// More flags to come.

func (p *flagParser) flag() string { _ = "STUB: not implemented"; return "" }

// We have "-ab", so return "-a" and keep "-b".

func (p *flagParser) value() string { _ = "STUB: not implemented"; return "" }

func (p *flagParser) args() []string { _ = "STUB: not implemented"; return nil }

type getopts struct {
	argidx  int
	runeidx int
}

func (g *getopts) next(optstr string, args []string) (opt rune, optarg string, done bool) {
	_ = "STUB: not implemented"
	return 0, "", false
}

// invalid option

// missing argument

// optStatusText returns a shell option's status text display
func (r *Runner) optStatusText(status bool) string { _ = "STUB: not implemented"; return "" }
