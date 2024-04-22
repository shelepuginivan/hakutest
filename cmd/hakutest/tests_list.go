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
	Run:     testsListCommand,
	Aliases: []string{"ls"},
}

func testsListCommand(cmd *cobra.Command, args []string) {
	for _, test := range test.NewService(app).GetTestList() {
		fmt.Println(test)
	}
}
