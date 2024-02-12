package main

import (
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
	"github.com/spf13/cobra"
)

func init() {
	testsCmd.AddCommand(testsRemoveCmd)
}

var testsRemoveCmd = &cobra.Command{
	Use:     "remove <test>",
	Short:   "Remove a test by its name",
	RunE:    testsRemoveCommand(test.NewService()),
	Aliases: []string{"rm"},
}

func testsRemoveCommand(t *test.TestService) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return t.Remove(args[0])
	}
}
