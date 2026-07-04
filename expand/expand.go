// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package expand

import (
	"io"
	"io/fs"
	"iter"
	"regexp"
	"strings"

	"mvdan.cc/sh/v3/syntax"
)

// A Config specifies details about how shell expansion should be performed. The
// zero value is a valid configuration.
type Config struct {
	// Env is used to get and set environment variables when performing
	// shell expansions. Some special parameters are also expanded via this
	// interface, such as:
	//
	//   * "#", "@", "*", "0"-"9" for the shell's parameters
	//   * "?", "$", "PPID" for the shell's status and process
	//   * "HOME foo" to retrieve user foo's home directory (if unset,
	//     os/user.Lookup will be used)
	//
	// If nil, there are no environment variables set. Use
	// ListEnviron(os.Environ()...) to use the system's environment
	// variables.
	Env Environ

	// CmdSubst expands a command substitution node, writing its standard
	// output to the provided [io.Writer].
	//
	// If nil, encountering a command substitution will result in an
	// UnexpectedCommandError.
	CmdSubst func(io.Writer, *syntax.CmdSubst) error

	// ProcSubst expands a process substitution node.
	ProcSubst func(*syntax.ProcSubst) (string, error)

	// TODO(v4): replace ReadDir with ReadDir2.

	// ReadDir is the older form of [ReadDir2], before io/fs.
	//
	// Deprecated: use ReadDir2 instead.
	ReadDir func(string) ([]fs.FileInfo, error)

	// ReadDir2 is used for file path globbing.
	// If nil, and [ReadDir] is nil as well, globbing is disabled.
	// Use [os.ReadDir] to use the filesystem directly.
	ReadDir2 func(string) ([]fs.DirEntry, error)

	// GlobStar corresponds to the shell option which allows globbing with "**".
	GlobStar bool

	// DotGlob corresponds to the shell option which allows filenames beginning
	// with a dot to be matched by a pattern which does not begin with a dot.
	DotGlob bool

	// NoCaseGlob corresponds to the shell option which causes case-insensitive
	// pattern matching in pathname expansion.
	NoCaseGlob bool

	// NullGlob corresponds to the shell option which allows globbing
	// patterns which match nothing to result in zero fields.
	NullGlob bool

	// NoUnset corresponds to the shell option which treats unset variables
	// as errors.
	NoUnset bool

	// ExtGlob corresponds to the shell option which allows using extended
	// pattern matching features when performing pathname expansion (globbing).
	ExtGlob bool

	bufferAlloc strings.Builder
	fieldAlloc  [4]fieldPart
	fieldsAlloc [4][]fieldPart

	ifs string
	// A pointer to a parameter expansion node, if we're inside one.
	// Necessary for ${LINENO}.
	curParam *syntax.ParamExp
}

// UnexpectedCommandError is returned if a command substitution is encountered
// when [Config.CmdSubst] is nil.
type UnexpectedCommandError struct {
	Node *syntax.CmdSubst
}

func (u UnexpectedCommandError) Error() string { _ = "STUB: not implemented"; return "" }

var zeroConfig = &Config{}

// TODO: note that prepareConfig is modifying the user's config in place,
// which doesn't feel right - we should make a copy.

func prepareConfig(cfg *Config) *Config { _ = "STUB: not implemented"; return nil }

func (cfg *Config) ifsRune(r rune) bool { _ = "STUB: not implemented"; return false }

func (cfg *Config) ifsJoin(strs []string) string { _ = "STUB: not implemented"; return "" }

func (cfg *Config) strBuilder() *strings.Builder { _ = "STUB: not implemented"; return nil }

func (cfg *Config) envGet(name string) string { _ = "STUB: not implemented"; return "" }

func (cfg *Config) envSet(name, value string) error { _ = "STUB: not implemented"; return nil }

// Literal expands a single shell word. It is similar to [Fields], but the result
// is a single string. This is the behavior when a word is used as the value in
// a shell variable assignment, for example.
//
// The config specifies shell expansion options; nil behaves the same as an
// empty config.
func Literal(cfg *Config, word *syntax.Word) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Document expands a single shell word as if it were a here-document body.
// It is similar to [Literal], but without brace expansion, tilde expansion, and
// globbing.
//
// The config specifies shell expansion options; nil behaves the same as an
// empty config.
func Document(cfg *Config, word *syntax.Word) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Pattern expands a single shell word as a pattern, using [pattern.QuoteMeta]
// on any non-quoted parts of the input word. The result can be used on
// [pattern.Regexp] directly.
//
// The config specifies shell expansion options; nil behaves the same as an
// empty config.
func Pattern(cfg *Config, word *syntax.Word) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Format expands a format string with a number of arguments, following the
// shell's format specifications. These include printf(1), among others.
//
// The resulting string is returned, along with the number of arguments used.
// Note that the resulting string may contain null bytes, for example
// if the format string used `\x00`. The caller should terminate the string
// at the first null byte if needed, such as when expanding for `$'foo\x00bar'`.
//
// The config specifies shell expansion options; nil behaves the same as an
// empty config.
func Format(cfg *Config, format string, args []string) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func formatInto(sb *strings.Builder, format string, args []string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// readDigits reads from 0 to max digits, either octal or
// hexadecimal.

// valid octal or hex char

// -1 since the outer loop does i++

// escaped

// bell

// backspace

// escape

// form feed

// new line

// carriage return

// horizontal tab

// vertical tab

// just the character

// if digits don't fit in 8 bits, 0xff via strconv

// can't error

// always as a single byte

// no escape sequence

// Passing in nil for args ensures that % format
// strings aren't processed; only escape sequences
// will be handled.

// if args == nil, we are not doing format
// arguments

func (cfg *Config) fieldJoin(parts []fieldPart) string { _ = "STUB: not implemented"; return "" }

// short-cut without a string copy

func (cfg *Config) escapedGlobField(parts []fieldPart) (escaped string, glob bool) {
	_ = "STUB: not implemented"
	return "", false
}

// only copy the string if it will be used

// Fields is a pre-iterators API which now wraps [FieldsSeq].
func Fields(cfg *Config, words ...*syntax.Word) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FieldsSeq expands a number of words as if they were arguments in a shell
// command. This includes brace expansion, tilde expansion, parameter expansion,
// command substitution, arithmetic expansion, quote removal, and globbing.
func FieldsSeq(cfg *Config, words ...*syntax.Word) iter.Seq2[string, error] {
	_ = "STUB: not implemented"
	return nil
}

// Note that globbing requires keeping a slice state, so it doesn't
// really benefit from using an iterator.

// We avoid [errors.As] as it allocates,
// and we know that [Config.glob] returns [pattern.Regexp] errors without wrapping.

// make a copy, since SplitBraces replaces the Parts slice

type fieldPart struct {
	val   string
	quote quoteLevel
}

type quoteLevel uint

const (
	quoteNone quoteLevel = iota
	quoteDouble
	quoteHeredoc
	quoteSingle
)

func (cfg *Config) wordField(wps []syntax.WordPart, ql quoteLevel) ([]fieldPart, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: return two separate fieldParts,
// like in wordFields?

// special chars

// write the special char, skipping the backslash

// TODO: why is this needed?

// cut the string if format included \x00

// Like how [Config.wordFields] deals with [syntax.ExtGlob],
// except that we allow these through even when [Config.ExtGlob]
// is false, as it only applies to pathname expansion.

func (cfg *Config) cmdSubst(cs *syntax.CmdSubst) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (cfg *Config) wordFields(wps []syntax.WordPart) ([][]fieldPart, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ending a field

// starting a new field

// ending a field without IFS

// cut the string if format included \x00

// We don't translate or interpret the pattern here in any way;
// that's done later when globbing takes place via [pattern.Regexp].
// Here, all we do is keep the extended globbing expression in string form.
//
// TODO(v4): perhaps the syntax parser should keep extended globbing expressions
// as plain literal strings, because a custom node is not particularly helpful.
// It's not like other globbing operators like `*` or `**` get their own nodes.

// quotedElemFields returns the list of elements resulting from a quoted
// parameter expansion that should be treated especially, like "${foo[@]}".
func (cfg *Config) quotedElemFields(pe *syntax.ParamExp) []string {
	_ = "STUB: not implemented"
	return nil
}

// "${!prefix@}"

// "${!prefix*}"

// "${!name[@]}"

// TODO: if an indexed array only has elements 0 and 10,
// we should not return all indices in between those.

// "${*}" or "${*:offset:length}"

// "${@}" or "${@:offset:length}"

// "${name[@]}"

// An unset variable expanded as "${name[@]}" produces
// zero fields, just like an empty array.

// "${name[*]}"

// sliceElems applies ${var:offset:length} slicing to a list of elements.
// When positional is true, $0 is prepended to the list before slicing.
// In bash, positional parameter offsets ($@ and $*) are 1-based and
// offset 0 includes $0 (the shell or script name). Negative offsets
// count from $# + 1, so $0 is reachable via large enough negative values.
func (cfg *Config) sliceElems(pe *syntax.ParamExp, elems []string, positional bool) []string {
	_ = "STUB: not implemented"
	return nil
}

func (cfg *Config) expandUser(field string, moreFields bool) (prefix, rest string) {
	_ = "STUB: not implemented"
	return "", ""
}

// No tilde prefix to expand, e.g. "foo".

// There is a tilde prefix, but followed by more fields, e.g. "~'foo'".
// We only proceed if an unquoted slash was found in this field, e.g. "~/'foo'".

// Current user; try via "HOME", otherwise fall back to the
// system's appropriate home dir env var. Don't use os/user, as
// that's overkill. We can't use [os.UserHomeDir], because we want
// to use cfg.Env, and we always want to check "HOME" first.

// Not the current user; try via "HOME <name>", otherwise fall back to
// os/user. There isn't a way to lookup user home dirs without cgo.

func findAllIndex(pat, name string, n int) [][]int { _ = "STUB: not implemented"; return nil }

var (
	rxGlobStar        = regexp.MustCompile(`^[^/.][^/]*$`)
	rxGlobStarDotGlob = regexp.MustCompile(`^[^/]*$`)
)

// pathJoin2 is a simpler version of [filepath.Join] without cleaning the result,
// since that's needed for globbing.
func pathJoin2(elem1, elem2 string) string { _ = "STUB: not implemented"; return "" }

// pathSplit splits a file path into its elements, retaining empty ones. Before
// splitting, slashes are replaced with [filepath.Separator], so that splitting
// Unix paths on Windows works as well.
func pathSplit(path string) []string { _ = "STUB: not implemented"; return nil }

func (cfg *Config) glob(base, pat string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// unix-like

// windows (for some reason it won't work without the
// trailing separator)

// TODO: as an optimization, we could do chunks of the path all at once,
// like doing a single stat for "/foo/bar" in "/foo/bar/*".

// TODO: Another optimization would be to reduce the number of ReadDir2 calls.
// For example, /foo/* can end up doing one duplicate call:
//
//    ReadDir2("/foo") to ensure that "/foo/" exists and only matches a directory
//    ReadDir2("/foo") glob "*"

// Keep around for debugging.
// log.Printf("matches %q part %d %q", matches, i, part)

// We can't use [Config.ReadDir2] on the parent and match the directory
// entry by name, because short paths on Windows break that.
// Our only option is to [Config.ReadDir2] on the directory entry itself,
// which can be wasteful if we only want to see if it exists,
// but at least it's correct in all scenarios.

// Unfortunately, [os.File.Readdir] on a regular file on
// Windows returns an error that satisfies [fs.ErrNotExist].
// Luckily, it returns a special "path not found" rather
// than the normal "file not found" for missing files,
// so we can use that knowledge to work around the bug.
// See https://github.com/golang/go/issues/46734.
// TODO: remove when the Go issue above is resolved.

// simply doesn't exist

// exists but not a directory

// Find all recursive matches for "**".
// Note that we need the results to be in depth-first order,
// and to avoid recursion, we use a slice as a stack.
// Since we pop from the back, we populate the stack backwards.

// "a/**" should match "a/ a/b a/b/cfg ...";
// note how the zero-match case there has a trailing separator.

// to reuse its capacity

// If dir is not a directory, we keep the stack as-is and continue.

// Note that the results need to be sorted.
// TODO: above we do a BFS; if we did a DFS, the matches would already be sorted.

// Remove any empty matches left behind from "**".

func (cfg *Config) globDir(base, dir string, matcher func(string) bool, wantDir bool, matches []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We still want to return matches, for the sake of reusing slices.

// No filtering.

// We need to know if the symlink points to a directory.
// This requires an extra syscall, as [Config.ReadDir] on the parent directory
// does not follow symlinks for each of the directory entries.
// ReadDir is somewhat wasteful here, as we only want its error result,
// but we could try to reuse its result as per the TODO in [Config.glob].

// Not a symlink nor a directory.

// ReadFields splits and returns n fields from s, like the "read" shell builtin.
// If raw is set, backslash escape sequences are not interpreted.
//
// The config specifies shell expansion options; nil behaves the same as an
// empty config.
func ReadFields(cfg *Config, s string, n int, raw bool) []string {
	_ = "STUB: not implemented"
	return nil
}

// include heading/trailing IFSs

// combine to max n fields
