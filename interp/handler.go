// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package interp

import (
	"context"
	"io"
	"io/fs"
	"os"
	"time"

	"mvdan.cc/sh/v3/expand"
	"mvdan.cc/sh/v3/syntax"
)

// HandlerCtx returns the [HandlerContext] value stored in ctx,
// which is used when calling handler functions.
// It panics if ctx has no HandlerContext stored.
func HandlerCtx(ctx context.Context) HandlerContext {
	_ = "STUB: not implemented"
	return *new(HandlerContext)
}

type handlerCtxKey struct{}

type handlerKind int

const (
	_                  handlerKind = iota
	handlerKindExec                // [ExecHandlerFunc]
	handlerKindCall                // [CallHandlerFunc]
	handlerKindOpen                // [OpenHandlerFunc]
	handlerKindReadDir             // [ReadDirHandlerFunc2]
)

// HandlerContext is the data passed to all the handler functions via [context.WithValue].
// It contains some of the current state of the [Runner].
type HandlerContext struct {
	runner *Runner // for internal use only, e.g. [HandlerContext.Builtin]

	// kind records which type of handler this context was built for.
	kind handlerKind

	// Env is a read-only version of the interpreter's environment,
	// including environment variables, global variables, and local function
	// variables.
	Env expand.Environ

	// Dir is the interpreter's current directory.
	Dir string

	// Pos is the source position which relates to the operation,
	// such as a [syntax.CallExpr] when calling an [ExecHandlerFunc].
	// It may be invalid if the operation has no relevant position information.
	Pos syntax.Pos

	// TODO(v4): use an os.File for stdin below directly.

	// Stdin is the interpreter's current standard input reader.
	// It is always an [*os.File], but the type here remains an [io.Reader]
	// due to backwards compatibility.
	Stdin io.Reader
	// Stdout is the interpreter's current standard output writer.
	Stdout io.Writer
	// Stderr is the interpreter's current standard error writer.
	Stderr io.Writer
}

// CallHandlerFunc is a handler which runs on every [syntax.CallExpr].
// It is called once variable assignments and field expansion have occurred.
// The context includes a [HandlerContext] value.
//
// The call's arguments are replaced by what the handler returns,
// and then the call is executed by the Runner as usual.
// The args slice is never empty.
// At this time, returning an empty slice without an error is not supported.
//
// This handler is similar to [ExecHandlerFunc], but has two major differences:
//
// First, it runs for all simple commands, including function calls and builtins.
//
// Second, it is not expected to execute the simple command, but instead to
// allow running custom code which allows replacing the argument list.
// Shell builtins touch on many internals of the Runner, after all.
//
// Returning a non-nil error will halt the [Runner] and will be returned via the API.
type CallHandlerFunc func(ctx context.Context, args []string) ([]string, error)

// TODO: consistently treat handler errors as non-fatal by default,
// but have an interface or API to specify fatal errors which should make
// the shell exit with a particular status code.

// ExecHandlerFunc is a handler which executes simple commands.
// It is called for all [syntax.CallExpr] nodes
// where the first argument is neither a declared function nor a builtin.
// The args slice is never empty.
// The context includes a [HandlerContext] value.
//
// Returning a nil error means a zero exit status.
// Other exit statuses can be set by returning or wrapping a [NewExitStatus] error,
// and such an error is returned via the API if it is the last statement executed.
// Any other error will halt the [Runner] and will be returned via the API.
type ExecHandlerFunc func(ctx context.Context, args []string) error

// DefaultExecHandler returns the [ExecHandlerFunc] used by default.
// It finds binaries in PATH and executes them.
// When context is cancelled, an interrupt signal is sent to running processes.
// killTimeout is a duration to wait before sending the kill signal.
// A negative value means that a kill signal will be sent immediately.
//
// On Windows, the kill signal is always sent immediately,
// because Go doesn't currently support sending Interrupt on Windows.
// [Runner] defaults to a killTimeout of 2 seconds.
func DefaultExecHandler(killTimeout time.Duration) ExecHandlerFunc {
	_ = "STUB: not implemented"
	return *new(ExecHandlerFunc)
}

// TODO: don't sleep in this goroutine if the program
// stops itself with the interrupt above.

// Windows and Plan9 do not have support for [syscall.WaitStatus]
// with methods like Signaled and Signal, so for those, [waitStatus] is a no-op.
// Note: [waitStatus] is an alias [syscall.WaitStatus]

// did not start

func checkStat(dir, file string, checkExec bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func winHasExt(file string) bool { _ = "STUB: not implemented"; return false }

// findExecutable returns the path to an existing executable file.
func findExecutable(dir, file string, exts []string) (string, error) {
	_ = "STUB: not implemented"
	return "",

		// non-windows
		nil
}

// findFile returns the path to an existing file.
func findFile(dir, file string, _ []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// LookPath is deprecated; see [LookPathDir].
func LookPath(env expand.Environ, file string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// LookPathDir is similar to [os/exec.LookPath], with the difference that it uses the
// provided environment. env is used to fetch relevant environment variables
// such as PWD and PATH.
//
// If no error is returned, the returned path must be valid.
func LookPathDir(cwd string, env expand.Environ, file string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// findAny defines a function to pass to [lookPathDir].
type findAny = func(dir string, file string, exts []string) (string, error)

func lookPathDir(cwd string, env expand.Environ, file string, find findAny) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// otherwise "foo" won't be "./foo"

// scriptFromPathDir is similar to [LookPathDir], with the difference that it looks
// for both executable and non-executable files.
func scriptFromPathDir(cwd string, env expand.Environ, file string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func pathExts(env expand.Environ) []string { _ = "STUB: not implemented"; return nil }

// OpenHandlerFunc is a handler which opens files.
// It is called for all files that are opened directly by the shell,
// such as in redirects, except for named pipes created by process substitutions.
// The context includes a [HandlerContext] value.
// Files opened by executed programs are not included.
//
// The path parameter may be relative to the current directory,
// which can be fetched via [HandlerCtx].
//
// Use a return error of type [*os.PathError] to have the error printed to
// stderr and the exit status set to 1.
// Any other error will halt the [Runner] and will be returned via the API.
//
// Note that implementations which do not return [os.File] will cause
// extra files and goroutines for input redirections; see [StdIO].
type OpenHandlerFunc func(ctx context.Context, path string, flag int, perm os.FileMode) (io.ReadWriteCloser, error)

// TODO: paths passed to [OpenHandlerFunc] should be cleaned.

// DefaultOpenHandler returns the [OpenHandlerFunc] used by default.
// It uses [os.OpenFile] to open files.
//
// For the sake of portability, /dev/null opens NUL on Windows.
func DefaultOpenHandler() OpenHandlerFunc { _ = "STUB: not implemented"; return *new(OpenHandlerFunc) }

// Note that even though https://go.dev/issue/71752 was resolved for Windows,
// the workaround here seems to still be required for Wine as of 10.14.
// TODO(mvdan): Why? Is this Wine's fault?

// TODO(v4): if this is kept in v4, it most likely needs to use [io/fs.DirEntry] for efficiency

// ReadDirHandlerFunc is a handler which reads directories. It is called during
// shell globbing, if enabled.
//
// Deprecated: use [ReadDirHandlerFunc2], which uses [fs.DirEntry].
type ReadDirHandlerFunc func(ctx context.Context, path string) ([]fs.FileInfo, error)

// ReadDirHandlerFunc2 is a handler which reads directories. It is called during
// shell globbing, if enabled.
// The context includes a [HandlerContext] value.
type ReadDirHandlerFunc2 func(ctx context.Context, path string) ([]fs.DirEntry, error)

// DefaultReadDirHandler returns the [ReadDirHandlerFunc] used by default.
// It makes use of [ioutil.ReadDir].
func DefaultReadDirHandler() ReadDirHandlerFunc {
	_ = "STUB: not implemented"
	return *new(ReadDirHandlerFunc)
}

// DefaultReadDirHandler2 returns the [ReadDirHandlerFunc2] used by default.
// It uses [os.ReadDir].
func DefaultReadDirHandler2() ReadDirHandlerFunc2 {
	_ = "STUB: not implemented"
	return *new(ReadDirHandlerFunc2)
}

// StatHandlerFunc is a handler which gets a file's information.
// The context includes a [HandlerContext] value.
type StatHandlerFunc func(ctx context.Context, name string, followSymlinks bool) (fs.FileInfo, error)

// DefaultStatHandler returns the [StatHandlerFunc] used by default.
// It makes use of [os.Stat] and [os.Lstat], depending on followSymlinks.
func DefaultStatHandler() StatHandlerFunc { _ = "STUB: not implemented"; return *new(StatHandlerFunc) }
