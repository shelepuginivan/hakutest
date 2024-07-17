package term

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// CorrectCommand displays the error message explaining the command in the
// following format:
//
//	error: <message>
//
//	    <correct-command>
//
// followed by [os.Exit](1) call.
//
// command should be a format string, containing `%s` for provided corrections.
func CorrectCommand(message, command string, corrections ...string) {
	var formatted []any
	for _, cor := range corrections {
		formatted = append(formatted, color.GreenString(cor))
	}

	correctCommand := fmt.Sprintf(command, formatted...)

	color.New(color.FgRed, color.Bold).Printf("error: %s\n\n", message)
	fmt.Printf("    %s\n\n", correctCommand)

	os.Exit(1)
}

// ErrorMultiline displays provided error messages in the following format:
//
//	error: <message 1>
//	       <message 2>
//	       ...
//
// followed by [os.Exit](1) call.
func ErrorMultiline(messages ...string) {
	c := color.New(color.FgRed, color.Bold)

	for i, msg := range messages {
		if i == 0 {
			c.Printf("error: %s\n", msg)
		} else {
			c.Printf("       %s\n", msg)
		}
	}

	os.Exit(1)
}
