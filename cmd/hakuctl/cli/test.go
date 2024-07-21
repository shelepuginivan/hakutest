package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/shelepuginivan/hakutest/internal/pkg/term"
	"github.com/shelepuginivan/hakutest/pkg/test"
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
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			term.CorrectCommand(
				"please provide at least one test.",
				"hakuctl test delete %s",
				"\"My Test\"",
			)
		}

		fmt.Printf("Tests marked for deletion: %s.\n", color.New(color.Bold).Sprint(len(args)))

		proceed, err := term.YesNo("Proceed?", true)
		for err != nil {
			proceed, err = term.YesNo("Proceed?", true)
		}

		if !proceed {
			fmt.Println("Cancelled.")
			return
		}

		removed := test.DeleteMany(args...)
		fmt.Printf("Deleted tests: %s.\n", color.New(color.Bold).Sprint(removed))
	},
}

var testExportCmd = &cobra.Command{
	Use:   "export <test1> [test2, test3, ...] -o <output>",
	Short: "Export test files",
	Example: `  hakuctl export "My Test.json" -o "Documents/My Test.json"         # Export single test to a JSON file
  hakuctl export test.json another.json third.json -o tests.zip     # Export multiple tests to a ZIP archive
  hakuctl export test.json -o -                                     # Export single test and print it to standard out`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			term.CorrectCommand(
				"please provide a test to export.",
				"hakuctl test export %s",
				"\"My test.json\"",
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

		if len(args) == 1 {
			testName := args[0]
			err := test.WriteJSON(out, testName)
			if err != nil {
				term.Error("an error occurred during export.", err.Error())
			}
			return
		}

		if err := test.WriteZip(out, args...); err != nil {
			term.Error("an error occurred during export", err.Error())
		}
	},
}

var testImportCmd = &cobra.Command{
	Use:   "import <file>",
	Short: "Import test from a file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			term.CorrectCommand(
				"please provide a file to import.",
				"hakuctl test import %s",
				"\"/path/to/my/file.json\"",
			)
		}

		importPath := args[0]
		data, err := os.ReadFile(importPath)

		if err != nil {
			term.Error(
				fmt.Sprintf("cannot read \"%s\".", importPath),
				"does it exist?",
			)
		}

		if test.Import(data) != nil {
			term.Error(
				fmt.Sprintf("cannot import \"%s\".", importPath),
				"is this a valid test?",
			)
		}

		fmt.Printf("Successfully imported \"%s\".\n", importPath)
	},
}

var testListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available test files",
	Run: func(cmd *cobra.Command, args []string) {
		for _, t := range test.GetList() {
			fmt.Println(t)
		}
	},
}

var testSearchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Incremental search among tests",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			term.CorrectCommand(
				"please provide a search query.",
				"hakuctl test search %s",
				"\"My test\"",
			)
		}

		for _, t := range test.GetList() {
			if strings.HasPrefix(t, args[0]) {
				fmt.Println(t)
			}
		}
	},
}
