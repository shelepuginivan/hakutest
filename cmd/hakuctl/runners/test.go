// Package runners provides functions for the command line interface.
// Exported methods are intended to be used in [cobra.Command] instances.
package runners

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/shelepuginivan/hakutest/internal/pkg/term"
	"github.com/shelepuginivan/hakutest/pkg/test"
	"github.com/spf13/cobra"
)

// TestDelete is a runner for `test delete` subcommand.
func TestDelete(cmd *cobra.Command, args []string) {
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
}

// TestExport is a runner for `test export` subcommand.
func TestExport(cmd *cobra.Command, args []string) {
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
}

// TestImport is a runner for `test import` subcommand.
func TestImport(cmd *cobra.Command, args []string) {
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
}

// TestList is a runner for `test list` subcommand.
func TestList(cmd *cobra.Command, args []string) {
	for _, t := range test.GetList() {
		fmt.Println(t)
	}
}

// TestSearch is a runner for `test search subcommand`.
func TestSearch(cmd *cobra.Command, args []string) {
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
}
