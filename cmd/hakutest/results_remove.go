package main

import (
	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/spf13/cobra"
)

func init() {
	resultsCmd.AddCommand(resultsRemoveCmd)
}

var resultsRemoveCmd = &cobra.Command{
	Use:     "remove <results>",
	Short:   "Remove a test by its name",
	RunE:    resultsRemoveCommand(results.NewService()),
	Aliases: []string{"rm"},
}

func resultsRemoveCommand(r *results.ResultsService) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return r.Remove(args[0])
	}
}
