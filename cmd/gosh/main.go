// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

// gosh is a proof of concept shell built on top of [interp].
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"mvdan.cc/sh/v3/interp"
)

var command = flag.String("c", "", "command to be executed")

func main() {
	flag.Parse()
	err := runAll()
	var es interp.ExitStatus
	if errors.As(err, &es) {
		os.Exit(int(es))
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runAll() error { _ = "STUB: not implemented"; return nil }

func run(r *interp.Runner, reader io.Reader, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func runPath(r *interp.Runner, path string) error { _ = "STUB: not implemented"; return nil }

func runInteractive(r *interp.Runner, stdin io.Reader, stdout, stderr io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// stop at the first error
