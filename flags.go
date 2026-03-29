package main

import (
	"flag"
)

var flagStdout bool

func init() {
	flag.BoolVar(&flagStdout, "stdout", false, "write to stdout instead of a file")
}
