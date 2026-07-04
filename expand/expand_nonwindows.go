// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

//go:build !windows

package expand

func isWindowsErrPathNotFound(error) bool { _ = "STUB: not implemented"; return false }
