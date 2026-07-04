// Copyright (c) 2016, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

// shfmt formats shell programs.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"runtime/debug"

	"golang.org/x/term"
	"mvdan.cc/editorconfig"

	"mvdan.cc/sh/v3/fileutil"
	"mvdan.cc/sh/v3/syntax"
)

type boolStringValue string

func (b *boolStringValue) Set(val string) error { _ = "STUB: not implemented"; return nil }

func (b *boolStringValue) String() string { _ = "STUB: not implemented"; return "" }

func (*boolStringValue) IsBoolFlag() bool { _ = "STUB: not implemented"; return false }

func boolStringVar(p *string, name string, value string, usage string) {
	_ = "STUB: not implemented"
	return
}

func langVariantVar(p *syntax.LangVariant, name string, value syntax.LangVariant, usage string) {
	_ = "STUB: not implemented"
	return
}

type multiFlag[T any] struct {
	short, long string
	val         T
}

func flagVal[T any](short, long string, val T, register func(*T, string, T, string)) *multiFlag[T] {
	_ = "STUB: not implemented"
	return nil
}

var (
	// Generic flags.
	versionFlag = flagVal("", "version", false, flag.BoolVar)
	list        = flagVal("l", "list", "false", boolStringVar)
	write       = flagVal("w", "write", false, flag.BoolVar)
	diff        = flagVal("d", "diff", false, flag.BoolVar)
	applyIgnore = flagVal("", "apply-ignore", false, flag.BoolVar)
	filename    = flagVal("", "filename", "", flag.StringVar)

	// Parser flags.
	lang     = flagVal("ln", "language-dialect", syntax.LangAuto, langVariantVar)
	posix    = flagVal("p", "posix", false, flag.BoolVar)
	simplify = flagVal("s", "simplify", false, flag.BoolVar)
	// TODO: when promoting exp.recover to a stable flag, add it as an EditorConfig knob too, and perhaps rename to recover-errors
	expRecover = flagVal("", "exp.recover", 0, flag.IntVar)

	// Printer flags.
	indent      = flagVal("i", "indent", 0, flag.UintVar)
	binNext     = flagVal("bn", "binary-next-line", false, flag.BoolVar)
	caseIndent  = flagVal("ci", "case-indent", false, flag.BoolVar)
	spaceRedirs = flagVal("sr", "space-redirects", false, flag.BoolVar)
	keepPadding = flagVal("kp", "keep-padding", false, flag.BoolVar)
	funcNext    = flagVal("fn", "func-next-line", false, flag.BoolVar)
	minify      = flagVal("mn", "minify", false, flag.BoolVar)

	// Utility flags.
	find     = flagVal("f", "find", "false", boolStringVar)
	toJSON   = flagVal("tojson", "to-json", false, flag.BoolVar) // TODO(v4): remove "tojson" for consistency
	fromJSON = flagVal("", "from-json", false, flag.BoolVar)

	// useEditorConfig will be false if any parser or printer flags were used.
	useEditorConfig = true

	parser            *syntax.Parser
	printer           *syntax.Printer
	readBuf, writeBuf bytes.Buffer
	color             bool

	copyBuf = make([]byte, 32*1024)
)

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, `usage: shfmt [flags] [path ...]

shfmt formats shell programs. If the only argument is a dash ('-') or no
arguments are given, standard input will be used. If a given path is a
directory, all shell scripts found under that directory will be used.

  --version  show version and exit

  -l[=0], --list[=0]  error with a list of files whose formatting differs from shfmt;
                      paths are separated by a newline or a null character if -l=0
  -w,     --write     write result to file instead of stdout
  -d,     --diff      error with a diff when the formatting differs
  --apply-ignore      always apply EditorConfig ignore rules
  --filename str      provide a name for the standard input file

Parser options:

  -ln, --language-dialect str  bash/posix/mksh/bats/zsh, default "auto"
  -p,  --posix                 shorthand for -ln=posix
  -s,  --simplify              simplify the code

Printer options:

  -i,  --indent uint       0 for tabs (default), >0 for number of spaces
  -bn, --binary-next-line  binary ops like && and | may start a line
  -ci, --case-indent       switch cases will be indented
  -sr, --space-redirects   redirect operators will be followed by a space
  -kp, --keep-padding      keep column alignment paddings
  -fn, --func-next-line    function opening braces are placed on a separate line
  -mn, --minify             minify the code to reduce its size (implies -s)

Utilities:

  -f[=0], --find[=0]  recursively find all shell files and print the paths;
                      paths are separated by a newline or a null character if -f=0
  --to-json           print syntax tree to stdout as a typed JSON
  --from-json         read syntax tree from stdin as a typed JSON

Formatting options can also be read from EditorConfig files; see 'man shfmt'
for a detailed description of the tool's behavior.
For more information and to report bugs, see https://github.com/mvdan/sh.
`)
	}
	flag.Parse()

	if versionFlag.val {
		version := "(unknown)"
		if info, ok := debug.ReadBuildInfo(); ok {
			mod := &info.Main
			if mod.Replace != nil {
				mod = mod.Replace
			}
			version = mod.Version
		}
		fmt.Println(version)
		return
	}
	if posix.val && lang.val != syntax.LangAuto {
		fmt.Fprintf(os.Stderr, "-p and -ln=lang cannot coexist\n")
		os.Exit(1)
	}
	if list.val != "true" && list.val != "false" && list.val != "0" {
		fmt.Fprintf(os.Stderr, "only -l and -l=0 allowed\n")
		os.Exit(1)
	}
	if find.val != "true" && find.val != "false" && find.val != "0" {
		fmt.Fprintf(os.Stderr, "only -f and -f=0 allowed\n")
		os.Exit(1)
	}
	simplify.val = simplify.val || minify.val
	flag.Visit(func(f *flag.Flag) {
		// This list should be in sync with the grouping of parser and printer options
		// as shown by ./shfmt.1.scd.
		switch f.Name {
		case lang.short, lang.long,
			posix.short, posix.long,
			simplify.short, simplify.long,
			indent.short, indent.long,
			binNext.short, binNext.long,
			caseIndent.short, caseIndent.long,
			spaceRedirs.short, spaceRedirs.long,
			keepPadding.short, keepPadding.long,
			funcNext.short, funcNext.long,
			minify.short, minify.long:
			useEditorConfig = false
		}
	})
	parser = syntax.NewParser(syntax.KeepComments(true))
	printer = syntax.NewPrinter(syntax.Minify(minify.val))

	syntax.RecoverErrors(expRecover.val)(parser)

	if !useEditorConfig {
		if posix.val {
			// -p equals -ln=posix
			lang.val = syntax.LangPOSIX
		}

		syntax.Indent(indent.val)(printer)
		syntax.BinaryNextLine(binNext.val)(printer)
		syntax.SwitchCaseIndent(caseIndent.val)(printer)
		syntax.SpaceRedirects(spaceRedirs.val)(printer)
		syntax.KeepPadding(keepPadding.val)(printer)
		syntax.FunctionNextLine(funcNext.val)(printer)
	}

	// Decide whether or not to use color for the diff output,
	// as described in shfmt.1.scd.
	if os.Getenv("FORCE_COLOR") != "" {
		color = true
	} else if os.Getenv("NO_COLOR") != "" || os.Getenv("TERM") == "dumb" {
	} else if term.IsTerminal(int(os.Stdout.Fd())) {
		color = true
	}
	// TODO(v4): show the help text on zero arguments,
	// having the user run `shfmt -` if they want to format stdin.
	// Using a dash is more explicit, and new users can easily be
	// confused by `shfmt` seemingly hanging forever.
	if flag.NArg() == 0 || (flag.NArg() == 1 && flag.Arg(0) == "-") {
		name := "<standard input>"
		if toJSON.val {
			name = "" // the default is not useful there
		}
		if filename.val != "" {
			name = filename.val
		}
		if err := formatStdin(name); err != nil {
			if err != errFormattingDiffers {
				fmt.Fprintln(os.Stderr, err)
			}
			os.Exit(1)
		}
		return
	}
	if filename.val != "" {
		fmt.Fprintln(os.Stderr, "-filename can only be used with stdin")
		os.Exit(1)
	}
	if toJSON.val {
		fmt.Fprintln(os.Stderr, "--to-json can only be used with stdin")
		os.Exit(1)
	}
	status := 0
	for _, path := range flag.Args() {
		explicit := true
		if err := filepath.WalkDir(path, func(path string, entry fs.DirEntry, err error) error {
			defer func() { explicit = false }()
			if err != nil {
				return err
			}
			if entry.IsDir() && vcsDir.MatchString(entry.Name()) {
				return filepath.SkipDir
			}
			// If the path is not an explicit arg, or if --apply-ignore is set,
			// we find and apply EditorConfig ignore rules.
			if !explicit || applyIgnore.val {
				// We don't know the language variant at this point yet, as we are walking directories
				// and we first want to tell if we should skip a path entirely.
				//
				// TODO: Should the call to Find with the language name check "ignore" too, then?
				// Otherwise, a [[bash]] section with ignore=true is effectively never used.
				//
				// TODO: Should there be a way to explicitly turn off ignore rules when walking?
				// Perhaps swapping the default to --apply-ignore=auto and allowing --apply-ignore=false?
				// I don't imagine it's a particularly useful scenario for now.
				props, err := ecQuery.Find(path, []string{"shell"})
				if err != nil {
					return err
				}
				if props.Get("ignore") == "true" {
					if entry.IsDir() {
						return filepath.SkipDir
					} else {
						return nil
					}
				}
			}
			conf := fileutil.ConfIsScript
			// If the path is an explicit arg and it points to a symlink,
			// resolve that symlink so that we don't try to format non-regular files.
			if explicit && entry.Type()&fs.ModeSymlink != 0 {
				info, err := os.Stat(path)
				if err != nil {
					return err
				}
				entry = fs.FileInfoToDirEntry(info)
			}
			// If the path is not an explicit arg, or it's not regular, or --find is set,
			// then we check for extensions and shebangs.
			if !explicit || !entry.Type().IsRegular() || find.val == "true" {
				conf = fileutil.CouldBeScript2(entry)
				if conf == fileutil.ConfNotScript {
					return nil
				}
			}
			err = formatPath(path, conf == fileutil.ConfIfShebang)
			if err == errFormattingDiffers {
				status = 1
				err = nil
			} else if err != nil {
				fmt.Fprintln(os.Stderr, err)
				status = 1
			}
			return nil
		}); err != nil {
			fmt.Fprintln(os.Stderr, err)
			status = 1
		}
	}
	os.Exit(status)
}

var vcsDir = regexp.MustCompile(`^\.(git|svn|hg)$`)

var errFormattingDiffers = fmt.Errorf("")

func formatStdin(name string) error { _ = "STUB: not implemented"; return nil }

// Mimic the logic from walkPath to apply the ignore rules.

// Fall back to bash.

func langFromFilename(name string) syntax.LangVariant {
	_ = "STUB: not implemented"
	// Detect shell config files, which typically have no extension nor shebang.
	// Note that these are not matched by [fileutil.CouldBeScript2],
	// because these files are typically not part of source code projects,
	// so finding these files when formatting entire directories is unnecessary.
	return *new(syntax.LangVariant)
}

// Same as the usual fallback, but we can avoid trying to read a shebang.

// fallback when none is found

// Note that ".sh" doesn't mean it's POSIX Shell for sure.
// A shebang in the file contents could say Bash or some other POSIX-like shell.

var ecQuery = editorconfig.Query{
	FileCache:   make(map[string]*editorconfig.File),
	RegexpCache: make(map[string]*regexp.Regexp),
}

func propsOptions(lang syntax.LangVariant, props editorconfig.Section) (_ syntax.LangVariant, validLang bool) {
	_ = "STUB: not implemented"
	// if shell_variant is set to a valid string, it will take precedence
	return *new(syntax.LangVariant), false
}

// TODO(v4): rename to case_indent for consistency with flags

// TODO(v4): rename to func_next_line for consistency with flags

// Note that --simplify is not actually a parser option, so we use a global var.
// Just like the CLI flags, minify=true implies simplify=true.

func formatPath(path string, checkShebang bool) error { _ = "STUB: not implemented"; return nil }

// only wanted the shebang for LangAuto

// too short to have a shebang

// some other read error

// not a shell script

// Fall back to bash.

func editorConfigLangs(l syntax.LangVariant) []string {
	_ = "STUB: not implemented"
	// All known shells match [[shell]].
	// As a special case, bash and the bash-like bats also match [[bash]],
	// and zsh also matches [[zsh]].
	// We can later consider others like [[mksh]] or [[posix-shell]],
	// just consider what list of languages the EditorConfig spec might eventually use.
	return nil
}

func formatBytes(src []byte, path string, fileLang syntax.LangVariant) error {
	_ = "STUB: not implemented"
	return nil
}

// Note that --simplify is treated as a parser option as it happens
// immediately after parsing, even if it's not a [syntax.ParserOption] today.

// must be standard input; fine to return
// TODO: change the default behavior to be compact,
// and allow using --to-json=pretty or --to-json=indent.

// TODO: support atomic writes on Windows?

// The first three lines are the header with the filenames, including --- and +++,
// and are marked in bold.

// the first three lines are bold

const (
	terminalGreen = "\u001b[32m"
	terminalRed   = "\u001b[31m"
	terminalCyan  = "\u001b[36m"
	terminalReset = "\u001b[0m"
	terminalBold  = "\u001b[1m"
)
