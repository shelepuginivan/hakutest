package cli

import (
	"github.com/shelepuginivan/hakutest/cmd/hakuctl/runners"
	"github.com/spf13/cobra"
)

func init() {
	// Flags for `hakuctl test` subcommands.
	testExportCmd.Flags().StringP("output", "o", "-", "Where to export files, - to export to stdout")

	// Add `hakuctl test` subcommands.
	testCmd.AddCommand(testDeleteCmd)
	testCmd.AddCommand(testExportCmd)
	testCmd.AddCommand(testImportCmd)
	testCmd.AddCommand(testListCmd)
	testCmd.AddCommand(testSearchCmd)

	// Add `hakuctl` subcommand.
	rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "test [command] [options]",
	Short: "Manage local test files",
}

var testDeleteCmd = &cobra.Command{
	Use:   "delete <test1> [test2, test3, ...]",
	Short: "Delete test files",
	Run:   runners.TestDelete,
}

var testExportCmd = &cobra.Command{
	Use:   "export <test1> [test2, test3, ...] -o <output>",
	Short: "Export test files",
	Example: `  hakuctl export "My Test.json" -o "Documents/My Test.json"         # Export single test to a JSON file
  hakuctl export test.json another.json third.json -o tests.zip     # Export multiple tests to a ZIP archive
  hakuctl export test.json -o -                                     # Export single test and print it to standard out`,
	Run: runners.TestExport,
}

var testImportCmd = &cobra.Command{
	Use:   "import <file>",
	Short: "Import test from a file",
	Run:   runners.TestImport,
}

var testListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available test files",
	Run:   runners.TestList,
}

var testSearchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Incremental search among tests",
	Run:   runners.TestSearch,
}
