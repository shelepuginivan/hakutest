package term

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// CorrectCommand writes the error message to STDERR explaining the command in
// the following format:
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

	color.New(color.FgRed, color.Bold).Fprintf(os.Stderr, "error: %s\n\n", message)
	fmt.Fprintf(os.Stderr, "  %s\n\n", correctCommand)

	os.Exit(1)
}

// ErrorMultiline writes provided error messages to STDERR in the following
// format:
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
			c.Fprintf(os.Stderr, "error: %s\n", msg)
		} else {
			c.Fprintf(os.Stderr, "       %s\n", msg)
		}
	}

	os.Exit(1)
}

// Warn writes provided warning message to STDERR in the following format:
//
//	warning: <message 1>
//	         <message 2>
//	         ...
func Warn(messages ...string) {
	c := color.New(color.FgYellow, color.Bold)

	for i, msg := range messages {
		if i == 0 {
			c.Fprintf(os.Stderr, "warning: %s\n", msg)
		} else {
			c.Fprintf(os.Stderr, "         %s\n", msg)
		}
	}
}
