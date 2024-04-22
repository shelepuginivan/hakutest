package main

import (
	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/spf13/cobra"
)

func resultsRemoveCommand(cmd *cobra.Command, args []string) error {
	return results.NewService(app).Remove(args[0])
}

var resultsRemoveCmd = &cobra.Command{
	Use:     "remove <results>",
	Short:   "Remove a test by its name",
	RunE:    resultsRemoveCommand,
	Aliases: []string{"rm"},
}
