package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	outputDir := outputDirectory()

	var when time.Time
	if unixTimestamp, ok := flagUnixTime.Get(); ok {
		when = time.Unix(unixTimestamp, 0)
	} else {
		when = time.Now()
	}

	unixTime := when.Unix()

	id, err := generateUUIDv7(when)
	if nil != err {
		fmt.Fprintf(os.Stderr, "ERROR: could not generate UUIDv7: %s\n", err)
		os.Exit(1)
	}

	content := citation(id, unixTime)

	if flagStdout {
		_, err = os.Stdout.Write(content)
		if nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: could not write to stdout: %s\n", err)
			os.Exit(1)
		}
		return
	}

	filename := id + ".citation"
	if "" != outputDir {
		filename = filepath.Join(outputDir, filename)
	}

	err = os.WriteFile(filename, content, 0664)
	if nil != err {
		fmt.Fprintf(os.Stderr, "ERROR: could not write file %q: %s\n", filename, err)
		os.Exit(1)
	}
}
