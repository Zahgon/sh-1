// Copyright (c) 2016, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package syntax

import (
	"bufio"
	"io"
	"text/tabwriter"
)

// PrinterOption is a function which can be passed to NewPrinter
// to alter its behavior. To apply option to existing Printer
// call it directly, for example KeepPadding(true)(printer).
type PrinterOption func(*Printer)

// Indent sets the number of spaces used for indentation. If set to 0,
// tabs will be used instead.
func Indent(spaces uint) PrinterOption { _ = "STUB: not implemented"; return *new(PrinterOption) }

// BinaryNextLine will make binary operators appear on the next line
// when a binary command, such as a pipe, spans multiple lines. A
// backslash will be used.
func BinaryNextLine(enabled bool) PrinterOption {
	_ = "STUB: not implemented"
	return *new(PrinterOption)
}

// SwitchCaseIndent will make switch cases be indented. As such, switch
// case bodies will be two levels deeper than the switch itself.
func SwitchCaseIndent(enabled bool) PrinterOption {
	_ = "STUB: not implemented"
	return *new(PrinterOption)
}

// TODO(v4): consider turning this into a "space all operators" option, to also
// allow foo=( bar baz ), (( x + y )), and so on.

// SpaceRedirects will put a space after most redirection operators. The
// exceptions are '>&', '<&', '>(', and '<('.
func SpaceRedirects(enabled bool) PrinterOption {
	_ = "STUB: not implemented"
	return *new(PrinterOption)
}

// KeepPadding will keep most nodes and tokens in the same column that
// they were in the original source. This allows the user to decide how
// to align and pad their code with spaces.
//
// Note that this feature is best-effort and will only keep the
// alignment stable, so it may need some human help the first time it is
// run.
//
// Deprecated: this formatting option is flawed and buggy, and often does
// not result in what the user wants when the code gets complex enough.
// The next major version, v4, will remove this feature entirely.
// See: https://github.com/mvdan/sh/issues/658
func KeepPadding(enabled bool) PrinterOption { _ = "STUB: not implemented"; return *new(PrinterOption) }

// Enable the flag, and set up the writer wrapper.

// Ensure we reset the state to that of NewPrinter.

// Minify will print programs in a way to save the most bytes possible.
// For example, indentation and comments are skipped, and extra
// whitespace is avoided when possible.
func Minify(enabled bool) PrinterOption { _ = "STUB: not implemented"; return *new(PrinterOption) }

// SingleLine will attempt to print programs in one line. For example, lists of
// commands or nested blocks do not use newlines in this mode. Note that some
// newlines must still appear, such as those following comments or around
// here-documents.
//
// Print's trailing newline when given a [*File] is not affected by this option.
func SingleLine(enabled bool) PrinterOption { _ = "STUB: not implemented"; return *new(PrinterOption) }

// FunctionNextLine will place a function's opening braces on the next line.
func FunctionNextLine(enabled bool) PrinterOption {
	_ = "STUB: not implemented"
	return *new(PrinterOption)
}

// NewPrinter allocates a new Printer and applies any number of options.
func NewPrinter(opts ...PrinterOption) *Printer { _ = "STUB: not implemented"; return nil }

// Print "pretty-prints" the given syntax tree node to the given writer. Writes
// to w are buffered.
//
// The node types supported at the moment are [*File], [*Stmt], [*Word], [*Assign], any
// [Command] node, and any WordPart node. A trailing newline will only be printed
// when a [*File] is used.
func (p *Printer) Print(w io.Writer, node Node) error { _ = "STUB: not implemented"; return nil }

// TODO: consider adding a raw mode to skip the tab writer, much like in
// go/printer.

// indenting with tabs

// indenting with spaces

// flush the writers

type bufWriter interface {
	Write([]byte) (int, error)
	WriteString(string) (int, error)
	WriteByte(byte) error
	Reset(io.Writer)
	Flush() error
}

type colCounter struct {
	*bufio.Writer
	column    int
	lineStart bool
}

func (c *colCounter) addByte(b byte) { _ = "STUB: not implemented"; return }

func (c *colCounter) WriteByte(b byte) error { _ = "STUB: not implemented"; return nil }

func (c *colCounter) WriteString(s string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *colCounter) Reset(w io.Writer) { _ = "STUB: not implemented"; return }

// Printer holds the internal state of the printing mechanism of a
// program.
type Printer struct {
	w         bufWriter
	tabWriter *tabwriter.Writer
	cols      colCounter // used for [KeepPadding]

	indentSpaces   uint
	binNextLine    bool
	swtCaseIndent  bool
	spaceRedirects bool
	keepPadding    bool
	minify         bool
	singleLine     bool
	funcNextLine   bool

	wantSpace wantSpaceState // whether space is required or has been written

	wantNewline bool // newline is wanted for pretty-printing; ignored by singleLine; ignored by singleLine
	mustNewline bool // newline is required to keep shell syntax valid
	wroteSemi   bool // wrote ';' for the current statement

	// pendingComments are any comments in the current line or statement
	// that we have yet to print. This is useful because that way, we can
	// ensure that all comments are written immediately before a newline.
	// Otherwise, in some edge cases we might wrongly place words after a
	// comment in the same line, breaking programs.
	pendingComments []Comment

	// firstLine means we are still writing the first line
	firstLine bool
	// line is the current line number
	line uint

	// lastLevel is the last level of indentation that was used.
	lastLevel uint
	// level is the current level of indentation.
	level uint
	// levelIncs records which indentation level increments actually
	// took place, to revert them once their section ends.
	levelIncs []bool

	nestedBinary bool

	// pendingHdocs is the list of pending heredocs to write.
	pendingHdocs []*Redirect

	// used when printing <<- heredocs with tab indentation
	tabsPrinter *Printer
}

func (p *Printer) reset() { _ = "STUB: not implemented"; return }

// minification uses its own newline logic

func (p *Printer) spaces(n uint) { _ = "STUB: not implemented"; return }

func (p *Printer) space() { _ = "STUB: not implemented"; return }

func (p *Printer) spacePad(pos Pos) { _ = "STUB: not implemented"; return }

// Never add padding at the start of a line unless we are indenting
// with spaces, since this may result in mixing of spaces and tabs.

// wantsNewline reports whether we want to print at least one newline before
// printing a node at a given position. A zero position can be given to simply
// tell if we want a newline following what's just been printed.
func (p *Printer) wantsNewline(pos Pos, escapingNewline bool) bool {
	_ = "STUB: not implemented"

	// We must have a newline here.
	return false
}

// The newline is optional, and singleLine skips it.
// Don't skip if there are any pending comments,
// as that might move them further down to the wrong place.

// The newline is optional, and we want it via either wantNewline or via
// the position's line.

func (p *Printer) bslashNewl() { _ = "STUB: not implemented"; return }

func (p *Printer) spacedString(s string, pos Pos) { _ = "STUB: not implemented"; return }

func (p *Printer) spacedToken(s string, pos Pos) { _ = "STUB: not implemented"; return }

func (p *Printer) semiOrNewl(s string, pos Pos) { _ = "STUB: not implemented"; return }

func (p *Printer) writeLit(s string) {
	_ = "STUB: not implemented"
	// If p.tabWriter is nil, this is the nested printer being used to print
	// <<- heredoc bodies, so the parent printer will add the escape bytes
	// later.
	return
}

func (p *Printer) incLevel() { _ = "STUB: not implemented"; return }

func (p *Printer) decLevel() { _ = "STUB: not implemented"; return }

func (p *Printer) indent() { _ = "STUB: not implemented"; return }

// TODO(mvdan): add an indent call at the end of newline?

// newline prints one newline and advances p.line to pos.Line().
func (p *Printer) newline(pos Pos) { _ = "STUB: not implemented"; return }

func (p *Printer) advanceLine(line uint) { _ = "STUB: not implemented"; return }

func (p *Printer) flushHeredocs() { _ = "STUB: not implemented"; return }

// Reuse the last indentation level, as
// indentation levels are usually changed before
// newlines are printed along with their
// subsequent indentation characters.

// The options need to persist.

// Overwrite p.line, since printing r.Word again can set
// p.line to the beginning of the heredoc again.

// newline prints between zero and two newlines.
// If any newlines are printed, it advances p.line to pos.Line().
func (p *Printer) newlines(pos Pos) { _ = "STUB: not implemented"; return }

// no empty lines at the top

// preserve single empty lines

func (p *Printer) rightParen(pos Pos) { _ = "STUB: not implemented"; return }

// closingParen prints a closing parenthesis at closePos, separating it from a
// preceding closing parenthesis on the same line to mirror the `( (` spacing
// that startsWithLparen adds to the matching opening parenthesis.
func (p *Printer) closingParen(stmts []*Stmt, last []Comment, openPos, closePos Pos) {
	_ = "STUB: not implemented"
	return
}

func (p *Printer) semiRsrv(s string, pos Pos) { _ = "STUB: not implemented"; return }

func (p *Printer) flushComments() { _ = "STUB: not implemented"; return }

// Flush any pending heredocs first. Otherwise, the comments would
// become part of a heredoc body. flushHeredocs may print and consume
// an inline comment, so range over pendingComments only after flushing,
// not over a stale copy that would reprint it after the heredoc.

// We can't call any of the newline methods, as they call this
// function and we'd recurse forever.

// don't go back one line, which may happen in some edge cases

func (p *Printer) comments(comments ...Comment) { _ = "STUB: not implemented"; return }

func (p *Printer) wordParts(wps []WordPart, quoted bool) {
	_ = "STUB: not implemented"
	// We disallow unquoted escaped newlines between word parts below.
	// However, we want to allow a leading escaped newline for cases such as:
	//
	//	foo <<< \
	//	  "bar baz"
	return
}

// Keep escaped newlines separating word parts when quoted.
// Note that those escaped newlines don't cause indentaiton.
// When not quoted, we strip them out consistently,
// because attempting to keep them would prevent indentation.
// Can't use p.wantsNewline here, since this is only about
// escaped newlines.

func (p *Printer) wordPart(wp, next WordPart) { _ = "STUB: not implemented"; return }

// ${10}
// ${var}cont

// avoid conflict with << and others

func (p *Printer) dblQuoted(dq *DblQuoted) { _ = "STUB: not implemented"; return }

// Add any trailing escaped newlines.

func (p *Printer) wroteIndex(index ArithmExpr) bool { _ = "STUB: not implemented"; return false }

// Note that e.g. foo[1,3]=$bar in Zsh does not allow any spaces around the comma,
// as that breaks the assignment word.

func (p *Printer) paramExp(pe *ParamExp) {
	_ = "STUB: not implemented"
	// arr[x]
	return
}

// Note that Zsh supports ${${nested}} but not ${$nested},
// so we need to avoid that simplification here.

func (p *Printer) cmdSubst(cs *CmdSubst) { _ = "STUB: not implemented"; return }

// Special case: `# inline comment`

func (p *Printer) loop(loop Loop) { _ = "STUB: not implemented"; return }

func (p *Printer) arithmExpr(expr ArithmExpr, compact, spacePlusMinus bool) {
	_ = "STUB: not implemented"
	return
}

func (p *Printer) arithmExprRecurse(expr ArithmExpr, compact, spacePlusMinus bool) {
	_ = "STUB: not implemented"
	return
}

func (p *Printer) testExpr(expr TestExpr) {
	_ = "STUB: not implemented"
	// Multi-line test expressions don't need to escape newlines.
	return
}

func (p *Printer) testExprSameLine(expr TestExpr) { _ = "STUB: not implemented"; return }

func (p *Printer) word(w *Word) { _ = "STUB: not implemented"; return }

func (p *Printer) unquotedWord(w *Word) { _ = "STUB: not implemented"; return }

func (p *Printer) wordJoin(ws []*Word) { _ = "STUB: not implemented"; return }

func (p *Printer) casePatternJoin(pats []*Word) { _ = "STUB: not implemented"; return }

// Only valid situation for a literal 'esac' here is with a preceding left paran.

func (p *Printer) elemJoin(elems []*ArrayElem, last []Comment) { _ = "STUB: not implemented"; return }

// Multi-line array expressions don't need to escape newlines.

func (p *Printer) stmt(s *Stmt) { _ = "STUB: not implemented"; return }

func (p *Printer) printRedirsUntil(redirs []*Redirect, startRedirs int, pos Pos) int {
	_ = "STUB: not implemented"
	return 0
}

func (p *Printer) command(cmd Command, redirs []*Redirect) (startRedirs int) {
	_ = "STUB: not implemented"
	return 0
}

// avoid ; in an empty block

// Forbid "foo()\n{ bar; }"

// Add a space between nested parentheses if we're printing them in a single line,
// to avoid the ambiguity between `((` and `( (`.

// Zsh allows empty subshells, but prevent `()`
// from looking like `() { anon-func; }`.

// leave p.nestedBinary untouched

// Apparently "case x in; esac" is invalid shell.

// avoid ; directly after tokens like ;;

func (p *Printer) ifClause(ic *IfClause, elif bool) { _ = "STUB: not implemented"; return }

func (p *Printer) stmtList(stmts []*Stmt, last []Comment) { _ = "STUB: not implemented"; return }

// In singleLine mode, ensure we use semicolons between
// statements.

// Comments after the end of this command. Note that
// this includes "<<EOF # comment".

// Comments between the beginning of the statement and
// the end of the command.

// The rest of the comments are before the entire
// statement.

func (p *Printer) nestedStmts(stmts []*Stmt, last []Comment, closing Pos) {
	_ = "STUB: not implemented"
	return
}

// Force a newline if we find:
//     { stmt; stmt; }

// Force a newline if we find:
//     { stmt
//     }

// Force a newline if we find:
//     for i in a b # stmt
//     do foo; done

func (p *Printer) assigns(assigns []*Assign) { _ = "STUB: not implemented"; return }

// Ensure we don't use an escaped newline after '=',
// because that can result in indentation, thus
// splitting "foo=bar" into "foo= bar".

type wantSpaceState uint8

const (
	spaceNotRequired wantSpaceState = iota
	spaceRequired                   // we should generally print a space or a newline next
	spaceWritten                    // we have just written a space or newline
)

// extraIndenter ensures that all lines in a '<<-' heredoc body have at least
// baseIndent leading tabs. Those that had more tab indentation than the first
// heredoc line will keep that relative indentation.
type extraIndenter struct {
	bufWriter
	baseIndent int

	firstIndent int
	firstChange int
	curLine     []byte
}

func (e *extraIndenter) WriteByte(b byte) error { _ = "STUB: not implemented"; return nil }

// no tabs if this is an empty line, i.e. "\n"

// This is the first heredoc line we add extra indentation to.
// Keep track of how much we indented.

// This line did not have enough indentation; simply indent it
// like the first line.

// This line had plenty of indentation. Add the extra
// indentation that the first line had, for consistency.

func (e *extraIndenter) WriteString(s string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func startsWithLparen(node Node) bool { _ = "STUB: not implemented"; return false }

// keep ( (

// keep ( ((

func endsWithRparen(node Node) bool { _ = "STUB: not implemented"; return false }

// keep ) )

// keep )) )
