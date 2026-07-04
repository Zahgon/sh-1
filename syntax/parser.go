// Copyright (c) 2016, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package syntax

import (
	"fmt"
	"io"
	"iter"
)

// ParserOption is a function which can be passed to NewParser
// to alter its behavior. To apply option to existing Parser
// call it directly, for example KeepComments(true)(parser).
type ParserOption func(*Parser)

// KeepComments makes the parser parse comments and attach them to
// nodes, as opposed to discarding them.
func KeepComments(enabled bool) ParserOption { _ = "STUB: not implemented"; return *new(ParserOption) }

// LangVariant describes a shell language variant to use when tokenizing and
// parsing shell code. The zero value is [LangBash].
//
// This type implements [flag.Value] so that it can be used as a CLI flag.
type LangVariant int

// TODO(v4): the zero value should be left as an unset and invalid value.
// TODO(v4): the type should be uint32 now that we use this as a bitset;
// an unsigned integer is clearer, and being agnostic to uint size avoids issues.

const (
	// LangBash corresponds to the GNU Bash language, as described in its
	// manual at https://www.gnu.org/software/bash/manual/bash.html.
	//
	// We currently follow Bash version 5.2.
	//
	// Its string representation is "bash".
	LangBash LangVariant = 1 << iota

	// LangPOSIX corresponds to the POSIX Shell language, as described at
	// https://pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap02.html.
	//
	// Its string representation is "posix" or "sh".
	LangPOSIX

	// LangMirBSDKorn corresponds to the MirBSD Korn Shell, also known as
	// mksh, as described at http://www.mirbsd.org/htman/i386/man1/mksh.htm.
	// Note that it shares some features with Bash, due to the shared
	// ancestry that is ksh.
	//
	// We currently follow mksh version 59.
	//
	// Its string representation is "mksh".
	LangMirBSDKorn

	// LangBats corresponds to the Bash Automated Testing System language,
	// as described at https://github.com/bats-core/bats-core. Note that
	// it's just a small extension of the Bash language.
	//
	// Its string representation is "bats".
	LangBats

	// LangZsh corresponds to the Z shell, as described at https://www.zsh.org/.
	//
	// Note that its support in the syntax package is experimental and
	// incomplete for now. See https://github.com/mvdan/sh/issues/120.
	//
	// We currently follow Zsh version 5.9.
	//
	// Its string representation is "zsh".
	LangZsh

	// LangAuto corresponds to automatic language detection,
	// commonly used by end-user applications like shfmt,
	// which can guess a file's language variant given its filename or shebang.
	//
	// At this time, [Variant] does not support LangAuto.
	LangAuto

	// langBashLegacy is what [LangBash] used to be, when it was zero.
	// We still support it for the sake of backwards compatibility.
	langBashLegacy LangVariant = 0

	// langResolvedVariants contains all known variants except [LangAuto],
	// which is meant to resolve to another variant.
	langResolvedVariants = LangBash | LangPOSIX | LangMirBSDKorn | LangBats | LangZsh

	// langResolvedVariantsCount is langResolvedVariants.count() as a constant.
	// TODO: Can we compute this as a constant expression somehow?
	// For example, if we had log2, we could do log2(LangAuto).
	langResolvedVariantsCount = 5

	// langBashLike contains Bash plus all variants which are extensions of it.
	langBashLike = LangBash | LangBats
)

// Variant changes the shell language variant that the parser will
// accept.
//
// The passed language variant must be one of the constant values defined in
// this package.
func Variant(l LangVariant) ParserOption { _ = "STUB: not implemented"; return *new(ParserOption) }

func (l LangVariant) String() string { _ = "STUB: not implemented"; return "" }

func (l *LangVariant) Set(s string) error { _ = "STUB: not implemented"; return nil }

func (l LangVariant) in(l2 LangVariant) bool { _ = "STUB: not implemented"; return false }

func (l LangVariant) count() int { _ = "STUB: not implemented"; return 0 }

func (l LangVariant) index() int { _ = "STUB: not implemented"; return 0 }

func (l LangVariant) bits() iter.Seq[LangVariant] { _ = "STUB: not implemented"; return nil }

// StopAt configures the lexer to stop at an arbitrary word, treating it
// as if it were the end of the input. It can contain any characters
// except whitespace, and cannot be over four bytes in size.
//
// This can be useful to embed shell code within another language, as
// one can use a special word to mark the delimiters between the two.
//
// As a word, it will only apply when following whitespace or a
// separating token. For example, StopAt("$$") will act on the inputs
// "foo $$" and "foo;$$", but not on "foo '$$'".
//
// The match is done by prefix, so the example above will also act on
// "foo $$bar".
func StopAt(word string) ParserOption { _ = "STUB: not implemented"; return *new(ParserOption) }

// RecoverErrors allows the parser to skip up to a maximum number of
// errors in the given input on a best-effort basis.
// This can be useful to tab-complete an interactive shell prompt,
// or when providing diagnostics on slightly incomplete shell source.
//
// Currently, this only helps with mandatory tokens from the shell grammar
// which are not present in the input. They result in position fields
// or nodes whose position report [Pos.IsRecovered] as true.
//
// For example, given the input
//
//	(foo |
//
// the result will contain two recovered positions; first, the pipe requires
// a statement to follow, and as [Stmt.Pos] reports, the entire node is recovered.
// Second, the subshell needs to be closed, so [Subshell.Rparen] is recovered.
func RecoverErrors(maximum int) ParserOption { _ = "STUB: not implemented"; return *new(ParserOption) }

// NewParser allocates a new [Parser] and applies any number of options.
func NewParser(options ...ParserOption) *Parser { _ = "STUB: not implemented"; return nil }

// Parse reads and parses a shell program with an optional name. It
// returns the parsed program if no issues were encountered. Otherwise,
// an error is returned. Reads from r are buffered.
//
// Parse can be called more than once, but not concurrently. That is, a
// Parser can be reused once it is done working.
func (p *Parser) Parse(r io.Reader, name string) (*File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EOF immediately after heredoc word so no newline to
// trigger the parsing error.

// Stmts is a pre-iterators API which now wraps [Parser.StmtsSeq].
//
// Deprecated: use [Parser.StmtsSeq].
func (p *Parser) Stmts(r io.Reader, fn func(*Stmt) bool) error {
	_ = "STUB: not implemented"
	return nil
}

// StmtsSeq reads and parses statements one at a time via an iterator.
func (p *Parser) StmtsSeq(r io.Reader) iter.Seq2[*Stmt, error] {
	_ = "STUB: not implemented"
	return nil
}

// EOF immediately after heredoc word so no newline to
// trigger the parsing error.

// Yield any final error from the parser.

type wrappedReader struct {
	p  *Parser
	rd io.Reader

	lastLine    int64
	accumulated []*Stmt
	yield       func([]*Stmt, error) bool
}

func (w *wrappedReader) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	// If we lexed a newline for the first time, we just finished a line, so
	// we may need to give a callback for the edge cases below not covered
	// by [Parser.Stmts].
	return 0, nil
}

// Incomplete statement; call back to print "> ".

// Nothing was parsed; call back to print another "$ ".

// Interactive is a pre-iterators API which now wraps [Parser.InteractiveSeq].
//
// Deprecated: use [Parser.InteractiveSeq].
func (p *Parser) Interactive(r io.Reader, fn func([]*Stmt) bool) error {
	_ = "STUB: not implemented"
	return nil
}

// InteractiveSeq implements what is necessary to parse statements in an
// interactive shell. The parser will call the given function under two
// circumstances outlined below.
//
// If a line containing any number of statements is parsed, the function will be
// called with said statements.
//
// If a line ending in an incomplete statement is parsed, the function will be
// called with any fully parsed statements, and [Parser.Incomplete] will return true.
//
// One can imagine a simple interactive shell implementation as follows:
//
//	fmt.Fprintf(os.Stdout, "$ ")
//	parser.Interactive(os.Stdin, func(stmts []*syntax.Stmt) bool {
//		if parser.Incomplete() {
//			fmt.Fprintf(os.Stdout, "> ")
//			return true
//		}
//		run(stmts)
//		fmt.Fprintf(os.Stdout, "$ ")
//		return true
//	}
//
// If the callback function returns false, parsing is stopped and the function
// is not called again.
func (p *Parser) InteractiveSeq(r io.Reader) iter.Seq2[[]*Stmt, error] {
	_ = "STUB: not implemented"
	return nil
}

// If the caller wishes, they can continue in the presence of parse errors.
// TODO: does this even work? Write tests for it. This only came up

// We finished parsing a statement and we're at a newline token,
// so we finished fully parsing a number of statements. Call
// back to run the statements and print "$ ".

// The callback above would already print "$ ", so we
// don't want the subsequent wrappedReader.Read to cause
// another "$ " print thinking that nothing was parsed.

// Words is a pre-iterators API which now wraps [Parser.WordsSeq].
//
// Deprecated: use [Parser.WordsSeq].
func (p *Parser) Words(r io.Reader, fn func(*Word) bool) error {
	_ = "STUB: not implemented"
	return nil
}

// WordsSeq reads and parses a sequence of words alongside any error encountered.
//
// Newlines are skipped, meaning that multi-line input will work fine. If the
// parser encounters a token that isn't a word, such as a semicolon, an error
// will be returned.
//
// Note that the lexer doesn't currently tokenize spaces, so it may need to read
// a non-space byte such as a newline or a letter before finishing the parsing
// of a word. This will be fixed in the future.
func (p *Parser) WordsSeq(r io.Reader) iter.Seq2[*Word, error] {
	_ = "STUB: not implemented"
	return nil
}

// Document parses a single here-document word. That is, it parses the input as
// if they were lines following a <<EOF redirection.
//
// In practice, this is the same as parsing the input as if it were within
// double quotes, but without having to escape all double quote characters.
// Similarly, the here-document word parsed here cannot be ended by any
// delimiter other than reaching the end of the input.
func (p *Parser) Document(r io.Reader) (*Word, error) { _ = "STUB: not implemented"; return nil, nil }

// Arithmetic parses a single arithmetic expression. That is, as if the input
// were within the $(( and )) tokens.
func (p *Parser) Arithmetic(r io.Reader) (ArithmExpr, error) {
	_ = "STUB: not implemented"
	return *new(ArithmExpr), nil
}

// Parser holds the internal state of the parsing mechanism of a
// program.
type Parser struct {
	src io.Reader
	bs  []byte // current chunk of read bytes
	bsp uint   // offset within [Parser.bs] for the rune after [Parser.r]
	r   rune   // next rune; [utf8.RuneSelf] when it went past EOF, or we stopped
	w   int    // width of [Parser.r]

	f *File

	spaced bool // whether [Parser.tok] has whitespace on its left

	err     error // lexer/parser error
	readErr error // got a read error, but bytes left
	readEOF bool  // [Parser.src] already gave us an [io.EOF] error

	tok token  // current token
	val string // current value (valid if tok is _Lit*)

	// position of [Parser.r], to be converted to [Parser.pos] later
	offs, line, col int64

	pos Pos // position of tok

	quote   quoteState // current lexer state
	eqlOffs int        // position of '=' in [Parser.val] when [Parser.tok].isLit is true

	keepComments bool
	lang         LangVariant

	stopAt []byte

	recoveredErrors  int
	recoverErrorsMax int

	forbidNested bool

	// list of pending heredoc bodies
	buriedHdocs int
	heredocs    []*Redirect

	hdocStops [][]byte // stack of end words for open heredocs

	parsingDoc bool // true if using [Parser.Document]

	// openNodes tracks how many entire statements or words we're currently parsing.
	// A non-zero number means that we require certain tokens or words before
	// reaching EOF, used for [Parser.Incomplete].
	openNodes int
	// openBquotes is how many levels of backquotes are open at the moment.
	openBquotes int

	// lastBquoteEsc is how many times the last backquote token was escaped
	lastBquoteEsc int

	rxOpenParens int
	rxFirstPart  bool

	accComs []Comment
	curComs *[]Comment

	litBatch  []Lit
	wordBatch []wordAlloc

	readBuf [bufSize]byte
	litBuf  [bufSize]byte
	litBs   []byte
}

// Incomplete reports whether the parser needs more input bytes
// to finish properly parsing a statement or word.
//
// It is only safe to call while the parser is blocked on a read. For an example
// use case, see [Parser.Interactive].
func (p *Parser) Incomplete() bool {
	_ = "STUB: not implemented"
	// If there are any open nodes, we need to finish them.
	// If we're constructing a literal, we need to finish it.
	return false
}

const bufSize = 1 << 10

func (p *Parser) reset() { _ = "STUB: not implemented"; return }

// nextPos returns the position of the next rune, [Parser.r].
func (p *Parser) nextPos() Pos {
	_ = "STUB: not implemented"
	// Basic protection against offset overflow;
	// note that an offset of 0 is valid, so we leave the maximum.
	return *new(Pos)
}

func (p *Parser) lit(pos Pos, val string) *Lit { _ = "STUB: not implemented"; return nil }

type wordAlloc struct {
	word  Word
	parts [1]WordPart
}

func (p *Parser) wordAnyNumber() *Word { _ = "STUB: not implemented"; return nil }

func (p *Parser) wordOne(part WordPart) *Word { _ = "STUB: not implemented"; return nil }

func (p *Parser) call(w *Word) *CallExpr { _ = "STUB: not implemented"; return nil }

type quoteState uint32

const (
	// The initial state of the parser.
	noState quoteState = 1 << iota

	// Used when parsing parameter expansions; use with [Parser.rune],
	// [Parser.next] always returns [illegalTok].
	runeByRune

	// unquotedWordCont exists purely so that the '#' in $foo#bar does not
	// get parsed as a comment; it's a tiny variation on [noState].
	unquotedWordCont

	subCmd
	subCmdBckquo
	dblQuotes
	hdocWord
	hdocBody
	hdocBodyTabs
	arithmExpr
	arithmExprLet
	arithmExprCmd
	testExpr
	testExprRegexp
	switchCase
	paramExpArithm
	paramExpRepl
	paramExpExp
	arrayElems

	allKeepSpaces = runeByRune | paramExpRepl | dblQuotes | hdocBody |
		hdocBodyTabs | paramExpRepl | paramExpExp
	allRegTokens = noState | unquotedWordCont | subCmd | subCmdBckquo | hdocWord |
		switchCase | arrayElems | testExpr
	allArithmExpr = arithmExpr | arithmExprLet | arithmExprCmd | paramExpArithm
	allParamExp   = paramExpArithm | paramExpRepl | paramExpExp
)

type saveState struct {
	quote       quoteState
	buriedHdocs int
}

func (p *Parser) preNested(quote quoteState) (s saveState) {
	_ = "STUB: not implemented"
	return *new(saveState)
}

func (p *Parser) postNested(s saveState) { _ = "STUB: not implemented"; return }

func (p *Parser) unquotedWordBytes(w *Word) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (p *Parser) unquotedWordPart(buf []byte, wp WordPart, quotes bool) (_ []byte, quoted bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (p *Parser) doHeredocs() { _ = "STUB: not implemented"; return }

// Nothing do do; don't even issue a read.

// consume '\n', since we know p.tok == _Newl

func (p *Parser) got(tok token) bool { _ = "STUB: not implemented"; return false }

func (p *Parser) gotRsrv(val string) (Pos, bool) {
	_ = "STUB: not implemented"
	return *new(Pos), false
}

func (p *Parser) recoverError() bool { _ = "STUB: not implemented"; return false }

type noQuote string

func (s noQuote) Format(f fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (t token) Format(f fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// EOF, Lit and the others should not be quoted in error messages
// as they are not real shell syntax like `if` or `{`.

func (p *Parser) followErr(pos Pos, left, right any) { _ = "STUB: not implemented"; return }

func (p *Parser) followErrExp(pos Pos, left any) { _ = "STUB: not implemented"; return }

func (p *Parser) follow(lpos Pos, left string, tok token) { _ = "STUB: not implemented"; return }

func (p *Parser) followRsrv(lpos Pos, left, val string) Pos {
	_ = "STUB: not implemented"
	return *new(Pos)
}

func (p *Parser) followStmts(left string, lpos Pos, stops ...string) ([]*Stmt, []Comment) {
	_ = "STUB: not implemented"
	// Language variants disallowing empty command lists:
	// * [LangPOSIX]: "A list is a sequence of one or more AND-OR lists...".
	// * [LangBash]: "A list is a sequence of one or more pipelines..."
	//
	// Language variants allowing empty command lists:
	// * [LangZsh]: "A list is a sequence of zero or more sublists...".
	// * [LangMirBSDKorn]: "Lists of commands can be created by separating pipelines...";
	//   note that the man page is not explicit, but the shell clearly allows e.g. `{ }`.
	return nil, nil
}

// allow an empty list

// allow an empty list

func (p *Parser) followWordTok(tok token, pos Pos) *Word { _ = "STUB: not implemented"; return nil }

func (p *Parser) stmtEnd(n Node, start, end string) Pos {
	_ = "STUB: not implemented"
	return *new(Pos)
}

func (p *Parser) quoteErr(lpos Pos, quote token) { _ = "STUB: not implemented"; return }

func (p *Parser) matchingErr(lpos Pos, left, right token) { _ = "STUB: not implemented"; return }

func (p *Parser) matched(lpos Pos, left, right token) Pos {
	_ = "STUB: not implemented"
	return *new(Pos)
}

func (p *Parser) errPass(err error) { _ = "STUB: not implemented"; return }

// IsIncomplete reports whether a Parser error could have been avoided with
// extra input bytes. For example, if an [io.EOF] was encountered while there was
// an unclosed quote or parenthesis.
func IsIncomplete(err error) bool { _ = "STUB: not implemented"; return false }

// TODO: probably redo with a [LangVariant] argument.
// Perhaps offer an iterator version as well.

// IsKeyword returns true if the given word is a language keyword
// in POSIX Shell or Bash.
func IsKeyword(word string) bool {
	_ = "STUB: not implemented"
	// This list has been copied from the bash 5.1 source code, file y.tab.c +4460
	// TODO: should we include entries for zsh here? e.g. "{}", "repeat", "always", ...
	return false
}

// only if COND_COMMAND is defined
// only if COND_COMMAND is defined

// only if COPROCESS_SUPPORT is defined

// only if SELECT_COMMAND is defined

// only if COMMAND_TIMING is defined

// ParseError represents an error found when parsing a source file, from which
// the parser cannot recover.
type ParseError struct {
	Filename string
	Pos      Pos
	Text     string

	Incomplete bool
}

func (e ParseError) Error() string { _ = "STUB: not implemented"; return "" }

// LangError is returned when the parser encounters code that is only valid in
// other shell language variants. The error includes what feature is not present
// in the current language variant, and what languages support it.
type LangError struct {
	Filename string
	Pos      Pos

	// TODO: consider replacing the Langs slice with a bitset.

	// Feature briefly describes which language feature caused the error.
	Feature string
	// Langs lists some of the language variants which support the feature.
	Langs []LangVariant
	// LangUsed is the language variant used which led to the error.
	LangUsed LangVariant
}

func (e LangError) Error() string { _ = "STUB: not implemented"; return "" }

func (p *Parser) posErr(pos Pos, format string, args ...any) {
	_ = "STUB: not implemented"
	//	for i, arg := range args {
	//		if arg, ok := arg.(fmt.Stringer); ok && arg != _EOF {
	//			args[i] = quotedToken(arg)
	//		}
	//	}
	return
}

func (p *Parser) curErr(format string, args ...any) { _ = "STUB: not implemented"; return }

func (p *Parser) checkLang(pos Pos, langSet LangVariant, format string, a ...any) {
	_ = "STUB: not implemented"
	return
}

// If we're reporting an error because a feature is for bash-like funcs,
// just mention "bash" rather than "bash/bats" for the sake of clarity.

func (p *Parser) stmts(yield func(*Stmt, error) bool, stops ...string) {
	_ = "STUB: not implemented"
	return
}

func (p *Parser) stmtList(stops ...string) ([]*Stmt, []Comment) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Split the comments, so that any aligned with an opening token
// get attached to it. For example:
//
//     if foo; then
//         # inside the body
//     # document the else
//     else
//     fi
// TODO(mvdan): look into deduplicating this with similar logic
// in caseItems.

// keep last nil if empty

func (p *Parser) invalidStmtStart() { _ = "STUB: not implemented"; return }

func (p *Parser) getWord() *Word { _ = "STUB: not implemented"; return nil }

func (p *Parser) getLit() *Lit { _ = "STUB: not implemented"; return nil }

func (p *Parser) wordParts(wps []WordPart) []WordPart { _ = "STUB: not implemented"; return nil }

// normalize empty lists into nil

func (p *Parser) ensureNoNested(pos Pos) { _ = "STUB: not implemented"; return }

func (p *Parser) wordPart() WordPart { _ = "STUB: not implemented"; return *new(WordPart) }

// don't tokenize '|'

// was not actually a parameter expansion, like: "foo$"

// p.tok == dblQuote, as "foo$" puts $ in the lit

// The lexer didn't call p.rune for us, so that it could have
// the right p.openBquotes to properly handle backslashes.

// e.g. found ` before the nested backquote \` was closed.

// Like above, the lexer didn't call p.rune for us.

// Zsh glob qualifier like *(N) or .(:a); the only case where
// ( immediately after a word is not a glob qualifier is ()
// for a function declaration, which the parser handles earlier.

// we can only get here due to EOF

func (p *Parser) cmdSubst() *CmdSubst { _ = "STUB: not implemented"; return nil }

func (p *Parser) dblQuoted() *DblQuoted { _ = "STUB: not implemented"; return nil }

// paramExp parses a short or full parameter expansion, depending on whether
// [Parser.tok] is [dollar] or [dollBrace]. It returns nil if a [dollar] token
// does not form a valid parameter expansion, in which case it should be parsed
// as a literal.
func (p *Parser) paramExp() *ParamExp { _ = "STUB: not implemented"; return nil }

// [ParamExp.Short] means we are parsing $exp rather than ${exp}.

// For now, for simplicity, we parse flags as just a literal.
// In the future, parsing as a word is better for cases like
// `${(ps.$sep.)val}`.

// we can only get here due to EOF

// Zsh-only prefixes that change how the parameter is expanded.
// They may appear in any combination, like ${=^name}.
// Doubling the rune (${==a}, ${~~a}, ${^^a}) forces the option off.

// For the short form, only treat as a prefix if followed by something
// that could start a parameter name or another zsh prefix.

// consume the first of the doubled pair

// Prefixes, like ${#name} to get the length of a variable.
// Note that in Zsh, the short form like $#name is allowed too.

// Unlike the others, zsh has no $!foo prefix.

// just "$"

// In short mode, any indexing or suffixes is not allowed, and we don't require '}'.
// Zsh is an exception: $foo[1] and $foo[1,3] are valid. Note that $1[x] does not qualify.

// Index expressions like ${foo[1]}. Note that expansion suffixes can be combined,
// like ${foo[@]//replace/with}.

// In zsh some of these like ${@[-1]} or ${*[1,3]} work,
// so we don't do this sort of check at all.

// pattern search and replace

// slicing

// Need to use a different matched style so arithm errors
// get reported correctly

// upper/lower case

func (p *Parser) paramNameStart() bool { _ = "STUB: not implemented"; return false }

func (p *Parser) nestedParameterStart(pe *ParamExp) (left token, quotePos Pos) {
	_ = "STUB: not implemented"
	return *new(token), *new(Pos)
}

// xxx given that we overwrite p.tok below

// '('

func (p *Parser) paramExpParameter(pe *ParamExp) *ParamExp {
	_ = "STUB: not implemented"
	// Check for Zsh nested parameter expressions like ${(f)"$(foo)"}.
	return nil
}

// ${#${nested parameter}}

// ${#$(nested command)}

// dollar

// The parameter name itself, like $foo or $?.

// actually ${#-default}, not ${#-}; fix the ambiguity

// Note that $1a is equivalent to ${1}a, but ${1a} is not.
// POSIX Shell says the latter is unspecified behavior, so match Bash's behavior.

// just "$"

// Zsh allows omitting the parameter name, e.g. ${:-word}.

func (p *Parser) paramExpExp() *Expansion { _ = "STUB: not implemented"; return nil }

func (p *Parser) eitherIndex() ArithmExpr { _ = "STUB: not implemented"; return *new(ArithmExpr) }

func (p *Parser) zshSubFlags() *FlagsArithm {
	_ = "STUB: not implemented"

	// Lex flags as raw text, like paramExp does for ${(flags)...}.
	return nil
}

// Lex the argument as a raw pattern, stopping at ',' or ']',
// since zsh treats it as a pattern rather than an arithmetic expression.

func (p *Parser) stopToken() bool { _ = "STUB: not implemented"; return false }

func (p *Parser) backquoteEnd() bool { _ = "STUB: not implemented"; return false }

// ValidName returns whether val is a valid name as per the POSIX spec.
func ValidName(val string) bool { _ = "STUB: not implemented"; return false }

func numberLiteral[T string | []byte](val T) bool { _ = "STUB: not implemented"; return false }

func (p *Parser) hasValidIdent() bool { _ = "STUB: not implemented"; return false }

// a+=x

// *[i]=x

// a[i]=x

func (p *Parser) getAssign(needEqual bool) *Assign { _ = "STUB: not implemented"; return nil }

// foo=bar

// a+=b

// since we're not using the entire p.val

// foo[x]=bar

// hasValidIdent already checks p.r is '['

// zsh allows a[i]=(values...).
// assgnParen consumed both '=' and '(',
// so rewrite as leftParen for array parsing below.

// TODO: support [index]=[

func (p *Parser) peekRedir() bool { _ = "STUB: not implemented"; return false }

func (p *Parser) doRedirect(s *Stmt) { _ = "STUB: not implemented"; return }

func (p *Parser) getStmt(readEnd, binCmd, fnBody bool) *Stmt { _ = "STUB: not implemented"; return nil }

// instead of using recursion, iterate manually

// left associativity: in a list of BinaryCmds, the
// right recursion should only read a single element

func (p *Parser) gotStmtPipe(s *Stmt, binCmd bool) *Stmt { _ = "STUB: not implemented"; return nil }

// Zsh treats closing braces in a special way, allowing this.

// TODO(zsh): "repeat"

// TODO(zsh): { try-list } "always" { always-list }

// Note that mksh lacks this one.

// Note that mksh lacks this one.

// In zsh, ( after a word is a glob qualifier unless followed
// immediately by ), which is the func declaration syntax.

// no statement found

// instead of using recursion, iterate manually

// left associativity: in a list of BinaryCmds, the
// right recursion should only read a single element

// No need to check for LangPOSIX, as on that language
// we parse |& as two tokens.

// in "! x | y", the bang applies to the entire pipeline

func (p *Parser) subshell(s *Stmt) { _ = "STUB: not implemented"; return }

func (p *Parser) arithmExpCmd(s *Stmt) { _ = "STUB: not implemented"; return }

func (p *Parser) block(s *Stmt) { _ = "STUB: not implemented"; return }

func (p *Parser) ifClause(s *Stmt) { _ = "STUB: not implemented"; return }

// All the nested IfClauses share the same FiPos.

func (p *Parser) whileClause(s *Stmt, until bool) { _ = "STUB: not implemented"; return }

func (p *Parser) forClause(s *Stmt) { _ = "STUB: not implemented"; return }

func (p *Parser) loop(fpos Pos) Loop { _ = "STUB: not implemented"; return *new(Loop) }

func (p *Parser) wordIter(ftok string, fpos Pos) *WordIter { _ = "STUB: not implemented"; return nil }

func (p *Parser) selectClause(s *Stmt) { _ = "STUB: not implemented"; return }

func (p *Parser) caseClause(s *Stmt) { _ = "STUB: not implemented"; return }

func (p *Parser) caseItems(stop string) (items []*CaseItem) { _ = "STUB: not implemented"; return nil }

// Split the comments:
//
// case x in
// a)
//   foo
//   ;;
//   # comment for a
// # comment for b
// b)
//   [...]

func (p *Parser) testClause(s *Stmt) { _ = "STUB: not implemented"; return }

func (p *Parser) testExprBinary(pastAndOr bool) TestExpr {
	_ = "STUB: not implemented"
	return *new(TestExpr)
}

// TODO(mvdan): Using nested states within a regex will break in
// all sorts of ways. The better fix is likely to use a stop
// token, like we do with heredocs.

func (p *Parser) testExprUnary() TestExpr { _ = "STUB: not implemented"; return *new(TestExpr) }

// not available in mksh

// otherwise we'd return a typed nil above

func (p *Parser) declClause(s *Stmt) { _ = "STUB: not implemented"; return }

func isBashCompoundCommand(tok token, val string) bool { _ = "STUB: not implemented"; return false }

func (p *Parser) timeClause(s *Stmt) { _ = "STUB: not implemented"; return }

func (p *Parser) coprocClause(s *Stmt) { _ = "STUB: not implemented"; return }

// has no name

// name was in fact the stmt

// name was in fact the start of a call

func (p *Parser) letClause(s *Stmt) { _ = "STUB: not implemented"; return }

func (p *Parser) bashFuncDecl(s *Stmt) { _ = "STUB: not implemented"; return }

// avoid non-nil zero-length slices

// allowed in all variants

func (p *Parser) testDecl(s *Stmt) { _ = "STUB: not implemented"; return }

func (p *Parser) unexpectedInCallExpr(ce *CallExpr) {
	_ = "STUB: not implemented"
	// Note that we'll only keep the first error that happens.
	return
}

func (p *Parser) callExpr(s *Stmt, w *Word, assign bool) { _ = "STUB: not implemented"; return }

// Avoid failing later with the confusing "} can only be used to close a block".

// Zsh does not require a semicolon to close a block.

func (p *Parser) funcDecl(s *Stmt, pos Pos, long, withParens bool, names ...*Lit) {
	_ = "STUB: not implemented"
	return
}

// TODO: reject any body which isn't a compound command, like a quoted word
