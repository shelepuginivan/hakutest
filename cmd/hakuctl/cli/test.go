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
	testCmd.AddCommand(testDeleteCmd)
	testCmd.AddCommand(testImportCmd)
	testCmd.AddCommand(testListCmd)
	testCmd.AddCommand(testSearchCmd)

	rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "test [command] [options]",
	Short: "Manage local test files",
}

var testDeleteCmd = &cobra.Command{
	Use:   "delete <test1> [tests2, test3, ...]",
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
			term.ErrorMultiline(
				fmt.Sprintf("cannot read \"%s\".", importPath),
				"does it exist?",
			)
		}

		if test.Import(data) != nil {
			term.ErrorMultiline(
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
