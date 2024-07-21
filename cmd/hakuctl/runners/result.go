package runners

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/shelepuginivan/hakutest/internal/pkg/term"
	"github.com/shelepuginivan/hakutest/pkg/result"
	"github.com/shelepuginivan/hakutest/pkg/statistics"
	"github.com/spf13/cobra"
)

// ResultDelete is a runner for `result delete` subcommand.
func ResultDelete(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		term.CorrectCommand(
			"please provide at least one result:",
			"hakuctl result delete %s",
			"\"My Result\"",
		)
	}

	fmt.Printf("Results marked for deletion: %s.\n", color.New(color.Bold).Sprint(len(args)))

	proceed, err := term.YesNo("Proceed?", true)
	for err != nil {
		proceed, err = term.YesNo("Proceed?", true)
	}

	if !proceed {
		fmt.Println("Cancelled")
		return
	}

	removed := result.DeleteMany(args...)
	fmt.Printf("Deleted results: %s.\n", color.New(color.Bold).Sprint(removed))
}

// ResultExport is a runner for `result export` subcommand.
func ResultExport(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		term.CorrectCommand(
			"please provide result name:",
			"hakuctl result export %s",
			"\"My result\"",
		)
	}

	name := args[0]

	s, err := statistics.NewFromSaved(name)
	if err != nil {
		term.Error(
			fmt.Sprintf("result %s not found.", name),
			"does it exist?",
		)
	}

	out, err := term.ResolveOutput(cmd.Flags().GetString("output"))
	if err != nil {
		term.Error(
			"cannot write to specified output",
			err.Error(),
		)
	}
	defer out.Close()

	err = nil

	switch f, _ := cmd.Flags().GetString("format"); f {
	case statistics.FormatCSV:
		err = s.WriteCSV(out)
	case statistics.FormatJSON:
		err = s.WriteJSON(out)
	case statistics.FormatXLSX:
		err = s.WriteXLSX(out)
	default:
		term.Warn(
			fmt.Sprintf("invalid format %s", f),
			fmt.Sprintf("falling back to %s",
				statistics.FormatJSON),
		)
		err = s.WriteJSON(out)
	}

	if err != nil {
		term.Error("an error occurred during export", err.Error())
	}
}

// ResultList is a runner for `result list` subcommand.
func ResultList(cmd *cobra.Command, args []string) {
	for _, t := range result.AvailableResults() {
		fmt.Println(t)
	}
}

// ResultSearch is a runner for `result search` subcommand.
func ResultSearch(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		term.CorrectCommand(
			"please provide a search query:",
			"hakuctl result search %s",
			"\"My result\"",
		)
	}

	for _, t := range result.AvailableResults() {
		if strings.HasPrefix(t, args[0]) {
			fmt.Println(t)
		}
	}
}
