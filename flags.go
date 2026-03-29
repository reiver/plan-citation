package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	libopt "plan-citation/lib/opt"
)

var flagStdout bool
var flagUnixTime libopt.OptionalInt64

func init() {
	// If we are running inside of a Go test, don't do all this.
	if nil != flag.Lookup("test.v") || strings.HasSuffix(os.Args[0], ".test") {
		return
	}

	flag.BoolVar(&flagStdout, "stdout", false, "write to stdout instead of a file")
	flag.Var(&flagUnixTime, "unix-time", "unix timestamp to use (default: current time)")

	flag.Parse()
}

func outputDirectory() string {
	args := flag.Args()
	switch len(args) {
	case 0:
		// use current directory
	case 1:
		info, err := os.Stat(args[0])
		if nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: could not stat %q: %s\n", args[0], err)
			os.Exit(1)
		}
		if !info.IsDir() {
			fmt.Fprintf(os.Stderr, "ERROR: %q is not a directory\n", args[0])
			os.Exit(1)
		}
		return args[0]
	default:
		fmt.Fprintf(os.Stderr, "ERROR: expected at most one directory argument, but got %d\n", len(args))
		os.Exit(1)
	}

	return ""
}
