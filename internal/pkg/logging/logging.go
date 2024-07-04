package logging

import (
	"io"
	"log"
	"os"
)

// output returns a temporary file to which logs are written.
// If creating a temporary file fails, it fallbacks to `os.Stdout`.
func output() io.Writer {
	tmp, err := os.CreateTemp("", "*.hakutest.log")
	if err != nil {
		return os.Stdout
	}

	return tmp
}

var Output io.Writer

func init() {
	Output = output()

	log.SetFlags(0)
	log.SetOutput(Output)
}
