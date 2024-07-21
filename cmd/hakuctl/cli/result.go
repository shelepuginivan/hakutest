package cli

import (
	"github.com/shelepuginivan/hakutest/cmd/hakuctl/runners"
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
	Run:   runners.ResultDelete,
}

var resultExportCmd = &cobra.Command{
	Use:   "export <result> -o <output> -f [format]",
	Short: "Generate and export result statistics",
	Run:   runners.ResultExport,
}

var resultListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available results",
	Run:   runners.ResultList,
}

var resultSearchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Incremental search among results",
	Run:   runners.ResultSearch,
}
