package main

import (
	"fmt"

	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/spf13/cobra"
)

func init() {
	resultsCmd.AddCommand(resultsListCmd)
}

var resultsListCmd = &cobra.Command{
	Use:     "list",
	Short:   "list available results",
	Run:     resultsListCommand(results.NewService()),
	Aliases: []string{"ls"},
}

func resultsListCommand(r *results.ResultsService) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		for _, result := range r.GetResultsList() {
			fmt.Println(result)
		}
	}
}
