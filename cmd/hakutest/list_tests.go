package main

import (
	"fmt"

	"github.com/shelepuginivan/hakutest/internal/pkg/test"
	"github.com/spf13/cobra"
)

func init() {
	listCmd.AddCommand(listTestsCmd)
}

var listTestsCmd = &cobra.Command{
	Use:   "tests",
	Short: "list available tests",
	Run:   listTestsCommand(test.NewService()),
}

func listTestsCommand(t *test.TestService) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		for _, test := range t.GetTestList() {
			fmt.Println(test)
		}
	}
}
