// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package interp

import (
	"context"
	"io"
	"io/fs"
	"iter"
	"os"
	"time"

	"mvdan.cc/sh/v3/expand"
	"mvdan.cc/sh/v3/syntax"
)

const (
	// shellReplyPS3Var, or PS3, is a special variable in Bash used by the select command,
	// while the shell is awaiting for input. the default value is [shellDefaultPS3]
	shellReplyPS3Var = "PS3"
	// shellDefaultPS3, or #?, is PS3's default value
	shellDefaultPS3 = "#? "
	// shellReplyVar, or REPLY, is a special variable in Bash that is used to store the result of
	// the select command or of the read command, when no variable name is specified
	shellReplyVar = "REPLY"

	fifoNamePrefix = "sh-interp-"
)

func (r *Runner) fillExpandConfig(ctx context.Context) { _ = "STUB: not implemented"; return }

// nothing to do

// $(<file)

// subshells don't exit the parent shell

// surface fatal errors immediately

// nothing to do

// We can't atomically create a random unused temporary FIFO.
// Similar to [os.CreateTemp],
// keep trying new random paths until one does not exist.
// We use a uint64 because a uint32 easily runs into retries.

// TODO: note that `man bash` mentions that `wait` only waits for the last
// process substitution as long as it is $!; the logic here would mean we wait for all of them.

// Should only happen if we forgot a case above.

// subshells don't exit the parent shell

// catShortcutArg checks if a statement is of the form "$(<file)". The redirect
// word is returned if there's a match, and nil otherwise.
func catShortcutArg(stmt *syntax.Stmt) *syntax.Word { _ = "STUB: not implemented"; return nil }

func (r *Runner) updateExpandOpts() { _ = "STUB: not implemented"; return }

func (r *Runner) expandErr(err error) { _ = "STUB: not implemented"; return }

// TODO: These errors are treated as fatal by bash.
// Make the error type reflect that.

// other cases do not exit

func (r *Runner) arithm(expr syntax.ArithmExpr) int { _ = "STUB: not implemented"; return 0 }

func (r *Runner) fields(words ...*syntax.Word) []string { _ = "STUB: not implemented"; return nil }

func (r *Runner) literal(word *syntax.Word) string { _ = "STUB: not implemented"; return "" }

func (r *Runner) document(word *syntax.Word) string { _ = "STUB: not implemented"; return "" }

func (r *Runner) pattern(word *syntax.Word) string { _ = "STUB: not implemented"; return "" }

// expandEnviron exposes [Runner]'s variables to the expand package.
type expandEnv struct {
	r *Runner
}

var _ expand.WriteEnviron = expandEnv{}

func (e expandEnv) Get(name string) expand.Variable {
	_ = "STUB: not implemented"
	return *new(expand.Variable)
}

func (e expandEnv) Set(name string, vr expand.Variable) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: return any errors

func (e expandEnv) Each(fn func(name string, vr expand.Variable) bool) {
	_ = "STUB: not implemented"
	return
}

var todoPos syntax.Pos // for handlerCtx callers where we don't yet have a position

func (r *Runner) handlerCtx(ctx context.Context, kind handlerKind, pos syntax.Pos) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// do not leave hc.Stdin as a typed nil

func (r *Runner) out(s string) { _ = "STUB: not implemented"; return }

func (r *Runner) outf(format string, a ...any) { _ = "STUB: not implemented"; return }

func (r *Runner) errf(format string, a ...any) { _ = "STUB: not implemented"; return }

func (r *Runner) stop(ctx context.Context) bool {
	_ = "STUB: not implemented"
	// Some traps trigger on exit, so we do want those to run.
	return false
}

func (r *Runner) stmt(ctx context.Context, st *syntax.Stmt) { _ = "STUB: not implemented"; return }

// subshells don't exit the parent shell

func (r *Runner) stmtSync(ctx context.Context, st *syntax.Stmt) { _ = "STUB: not implemented"; return }

// If the "errexit" option is set and a command failed, exit the shell. Exceptions:
//
//   conditions (if <cond>, while <cond>, etc)
//   part of && or || lists; excluded via "else" above
//   preceded by !; excluded via "else" above

func (r *Runner) cmd(ctx context.Context, cm syntax.Command) { _ = "STUB: not implemented"; return }

// subshells don't exit the parent shell

// Use a new slice, to not modify the slice in the alias map.

// Here we have a naked "foo=bar", so if we inherited a local var from a parent
// function we want to signal that we are modifying the parent var rather than
// creating a new local var via "local foo=bar".
// TODO: there is likely a better way to do this.

// Strangely enough, it seems like Bash prints original
// source for arrays, but the expanded value otherwise.
// TODO: add test cases for x[i]=y and x+=y.

// should never happen

// If interpreting the last expansion like $(foo) failed,
// and the expansion and assignments otherwise succeeded,
// we need to surface that last exit code.

// Resolve any nameref so we can restore the original final value later on.

// Inline command vars are always exported.

// not being able to create a pipe is rare but critical

// subshells don't exit the parent shell

// surface fatal errors immediately

// for i; do ...

// for i in ...; do ...

// display menu

// no reply; try again

// execute commands until break or return is encountered

// TODO

// to preserve exit status code 2 for regex errors, etc

// "-f" or "-p" for query mode

// When used in a function, "declare" acts as "local"
// unless the "-g" option is used.

// declare -f name: print function definition.
// Bash silently returns exit 1 for missing functions.

// declare -p name: print variable with attributes.

// TODO: can we do these?

// Should only happen if we forgot a case above.

func (r *Runner) trapCallback(ctx context.Context, callback, name string) {
	_ = "STUB: not implemented"
	return
}

// nothing to do

// don't recurse, as that could lead to cycles

// TODO: do this parsing when "trap" is called?

// ignore errors in the callback

// traps on EXIT or ERR should not modify the result

func (r *Runner) flattenAssigns(args []*syntax.Assign) iter.Seq[*syntax.Assign] {
	_ = "STUB: not implemented"
	return nil
}

// Convert "declare $x" into "declare value".
// Don't use syntax.Parser here, as we only want the basic
// splitting by '='.

func match(pat, name string) bool { _ = "STUB: not implemented"; return false }

// TODO: report these errors

func elapsedString(d time.Duration, posix bool) string { _ = "STUB: not implemented"; return "" }

func (r *Runner) stmts(ctx context.Context, stmts []*syntax.Stmt) {
	_ = "STUB: not implemented"
	return
}

func (r *Runner) hdocReader(rd *syntax.Redirect) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We write to the pipe in a new goroutine,
// as pipe writes may block once the buffer gets full.
// We still construct and buffer the entire heredoc first,
// as doing it concurrently would lead to different semantics and be racy.

func (r *Runner) redir(ctx context.Context, rd *syntax.Redirect) (io.Closer, error) {
	_ = "STUB: not implemented"
	return *new(io.Closer), nil
}

// Note that the input redirects below always use stdin (0)
// because we don't support anything else right now.

// The default for the output redirects below.

// We write to the pipe in a new goroutine,
// as pipe writes may block once the buffer gets full.

// closing the output writer

// done further below

// closing the input file

func (r *Runner) loopStmtsBroken(ctx context.Context, stmts []*syntax.Stmt) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *Runner) call(ctx context.Context, pos syntax.Pos, args []string) {
	_ = "STUB: not implemented"
	return
}

// handler's custom fatal error

// stack them to support nested func calls

// Functions run in a nested scope.
// Note that [Runner.exec] below does something similar.

func (r *Runner) exec(ctx context.Context, pos syntax.Pos, args []string) {
	_ = "STUB: not implemented"
	return
}

func (r *Runner) open(ctx context.Context, path string, flags int, mode os.FileMode, print bool) (io.ReadWriteCloser, error) {
	_ = "STUB: not implemented"
	// If we are opening a FIFO temporary file created by the interpreter itself,
	// don't pass this along to the open handler as it will not work at all
	// unless [os.OpenFile] is used directly with it.
	// Matching by directory and basename prefix isn't perfect, but works.
	//
	// If we want FIFOs to use a handler in the future, they probably
	// need their own separate handler API matching Unix-like semantics.
	return *new(io.ReadWriteCloser), nil
}

// TODO: support wrapped PathError returned from openHandler.

// handler's custom fatal error

func (r *Runner) stat(ctx context.Context, name string) (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}

func (r *Runner) lstat(ctx context.Context, name string) (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}
