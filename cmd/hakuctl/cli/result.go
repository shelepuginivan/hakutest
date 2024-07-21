package cli

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/shelepuginivan/hakutest/internal/pkg/term"
	"github.com/shelepuginivan/hakutest/pkg/result"
	"github.com/shelepuginivan/hakutest/pkg/statistics"
	"github.com/spf13/cobra"
)

func init() {
	// Flags for `hakuctl result` subcommands.
	resultExportCmd.Flags().StringP(
		"format", "f", statistics.FormatJSON,
		"Statistics export format. Valid values are \"csv\", \"json\", and \"xlsx\"",
	)
	resultExportCmd.Flags().StringP(
		"output", "o", "-",
		"Where to export files, - to export to stdout",
	)

	// Add `hakuctl result` subcommands.
	resultCmd.AddCommand(resultDeleteCmd)
	resultCmd.AddCommand(resultExportCmd)
	resultCmd.AddCommand(resultListCmd)
	resultCmd.AddCommand(resultSearchCmd)

	// Add `hakuctl` subcommand.
	rootCmd.AddCommand(resultCmd)
}

var resultCmd = &cobra.Command{
	Use:   "result",
	Short: "Manage test results and statistics",
}

var resultDeleteCmd = &cobra.Command{
	Use:   "delete <result1> [result2, result3, ...]",
	Short: "Delete results",
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

var resultExportCmd = &cobra.Command{
	Use:   "export <result> -o <output> -f [format]",
	Short: "Generate and export result statistics",
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

var resultListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available results",
	Run: func(cmd *cobra.Command, args []string) {
		for _, t := range result.AvailableResults() {
			fmt.Println(t)
		}
	},
}

var resultSearchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Incremental search among results",
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}
