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
	RunE:    testsRemoveCommand,
	Aliases: []string{"rm"},
}

func testsRemoveCommand(cmd *cobra.Command, args []string) error {
	t := test.NewService(app)
	return t.Remove(args[0])
}
