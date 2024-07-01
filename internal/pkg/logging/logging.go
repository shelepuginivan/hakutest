// Package logging provides logging methods and helpers.
package logging

import (
	"io"
	"os"
)

// Output returns a temporary file to which logs are written.
// If creating a temporary file fails, it fallbacks to `os.Stdout`.
func Output() io.Writer {
	tmp, err := os.CreateTemp("", "*.hakutest.log")
	if err != nil {
		return os.Stdout
	}

	return tmp
}
