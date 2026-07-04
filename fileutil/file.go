// Copyright (c) 2016, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

// Package fileutil allows inspecting shell files, such as detecting whether a
// file may be shell or extracting its shebang.
package fileutil

import (
	"io/fs"
	"regexp"
)

var (
	shebangRe = regexp.MustCompile(`^#![ \t]*/(usr/)?bin/(env[ \t]+)?(sh|dash|bash|mksh|bats|zsh)(\s|$)`)
	extRe     = regexp.MustCompile(`\.(sh|bash|mksh|bats|zsh)$`)
)

// TODO: consider removing HasShebang in favor of Shebang in v4

// HasShebang reports whether bs begins with a valid shell shebang.
// It supports variations with /usr and env.
func HasShebang(bs []byte) bool { _ = "STUB: not implemented"; return false }

// Shebang parses a "#!" sequence from the beginning of the input bytes,
// and returns the shell that it points to.
//
// For instance, it returns "sh" for "#!/bin/sh",
// and "bash" for "#!/usr/bin/env bash".
func Shebang(bs []byte) string { _ = "STUB: not implemented"; return "" }

// ScriptConfidence defines how likely a file is to be a shell script,
// from complete certainty that it is not one to complete certainty that
// it is one.
type ScriptConfidence int

const (
	// ConfNotScript describes files which are definitely not shell scripts,
	// such as non-regular files or files with a non-shell extension.
	ConfNotScript ScriptConfidence = iota

	// ConfIfShebang describes files which might be shell scripts, depending
	// on the shebang line in the file's contents. Since [CouldBeScript] only
	// works on [fs.FileInfo], the answer in this case can't be final.
	ConfIfShebang

	// ConfIsScript describes files which are definitely shell scripts,
	// which are regular files with a valid shell extension.
	ConfIsScript
)

// CouldBeScript is a shortcut for CouldBeScript2(fs.FileInfoToDirEntry(info)).
//
// Deprecated: prefer [CouldBeScript2], which usually requires fewer syscalls.
func CouldBeScript(info fs.FileInfo) ScriptConfidence {
	_ = "STUB: not implemented"
	return *new(ScriptConfidence)
}

// CouldBeScript2 reports how likely a directory entry is to be a shell script.
// It discards directories and other non-regular files like symbolic links,
// filenames beginning with '.', and files with non-shell extensions.
func CouldBeScript2(entry fs.DirEntry) ScriptConfidence {
	_ = "STUB: not implemented"
	return *new(ScriptConfidence)
}

// '.' prefix (hidden file)

// dir, symlink, named pipes, etc

// shell extension

// non-shell extension

// no extension; read and look for a shebang
