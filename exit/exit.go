// A package to reduce the if err != nil blocks to one-liners
package exit

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Determine which io.Writer to use based on exitCode. All non-zero exitCodes
// return io.Stderr.
func assignWriter(exitCode int) io.Writer {
	var w io.Writer
	if exitCode != 0 {
		w = os.Stderr
	} else {
		w = os.Stdout
	}

	return w
}

// exit.If b is true
func If(b bool, msg string, exitCode int) {

	if b {
		w := assignWriter(exitCode)

		if msg != "" {
			fmt.Fprintf(w, "%s\n", msg)
		}

		os.Exit(exitCode)
	}
}

// exit.If b is true
func IfErr(b bool, err error, exitCode int) {

	if b {
		w := assignWriter(exitCode)
		fmt.Fprintf(w, "%v\n", err)
		os.Exit(exitCode)
	}
}

// exit.Unless b is true
func Unless(b bool, msg string, exitCode int) {

	if !b {
		w := assignWriter(exitCode)

		if msg != "" {
			fmt.Fprintf(w, "%s\n", msg)
		}

		os.Exit(exitCode)
	}
}

// exit.AndLogError will use the log package settings used by the caller app
// and exit with the specified exitCode
func AndLogError(err error, exitCode int) {
	log.Println(err)
	os.Exit(exitCode)
}

// exit.AndWriteError will send the error message to stderr and exit with
// the specified exitCode
func AndWriteError(err error, exitCode int) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(exitCode)
}

// SDG
