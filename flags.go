package main

import (
	"flag"
	"os"
	"strings"
)

var flagStdout bool

func init() {
	// If we are running inside of a Go test, don't do all thius.
	if nil != flag.Lookup("test.v") || strings.HasSuffix(os.Args[0], ".test") {
		return
	}

	flag.BoolVar(&flagStdout, "stdout", false, "write to stdout instead of a file")

	flag.Parse()
}
