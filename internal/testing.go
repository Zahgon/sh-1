// Copyright (c) 2026, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package internal

// TestMainSetup is used by the integration tests running shell scripts
// either via our interpreter or via real shells,
// to ensure a reasonably clean and consistent environment.
func TestMainSetup() {
	_ = "STUB: not implemented"
	// Set the locale to computer-friendly English and UTF-8.
	// Some systems like macOS miss C.UTF8, so fall back to the US English locale.
	return
}

// Bash prints the pwd after changing directories when CDPATH is set.

// These short names are commonly used as variables.
// Ensure they are unset as env vars.
// We can't easily remove names from $PATH,
// so do the next best thing: override each name with a failing script.
