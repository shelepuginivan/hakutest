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
	Short: "Manage test files.",
}

var testDeleteCmd = &cobra.Command{
	Use:   "delete <test1> [tests2, test3, ...]",
	Short: "Delete test files.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			color.Red("error: please provide at least one test.\n\n")
			fmt.Printf("    hakuctl test delete %s\n\n", color.GreenString("\"My test\""))
			os.Exit(1)
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
			color.Red("error: please provide a file to import.\n\n")
			fmt.Printf("    hakuctl test import %s\n\n", color.GreenString("\"/path/to/my/file.json\""))
			os.Exit(1)
		}

		importPath := args[0]
		data, err := os.ReadFile(importPath)

		if err != nil {
			color.Red("error: cannot read \"%s\".\n", importPath)
			color.Red("does it exist?\n")
			os.Exit(1)
		}

		if test.Import(data) != nil {
			color.Red("error: cannot import \"%s\".\n", importPath)
			color.Red("is this a valid test?\n")
			os.Exit(1)
		}

		fmt.Printf("Successfully imported \"%s\".\n", importPath)
	},
}

var testListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available test files.",
	Run: func(cmd *cobra.Command, args []string) {
		for _, t := range test.GetList() {
			fmt.Println(t)
		}
	},
}

var testSearchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Incremental search among tests.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			color.Red("error: please provide a search query.\n\n")
			fmt.Printf("    hakuctl test search %s\n\n", color.GreenString("\"My test\""))
			os.Exit(1)
		}

		for _, t := range test.GetList() {
			if strings.HasPrefix(t, args[0]) {
				fmt.Println(t)
			}
		}
	},
}
