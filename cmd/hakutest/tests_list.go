package main

import (
	"fmt"

	"github.com/shelepuginivan/hakutest/internal/pkg/test"
	"github.com/spf13/cobra"
)

func init() {
	testsCmd.AddCommand(testsListCmd)
}

var testsListCmd = &cobra.Command{
	Use:     "list",
	Short:   "list available tests",
	Run:     testsListCommand(test.NewService()),
	Aliases: []string{"ls"},
}

func testsListCommand(t *test.TestService) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		for _, test := range t.GetTestList() {
			fmt.Println(test)
		}
	}
}
