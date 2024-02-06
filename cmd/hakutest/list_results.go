package main

import (
	"fmt"

	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/spf13/cobra"
)

func init() {
	listCmd.AddCommand(listResultsCmd)
}

var listResultsCmd = &cobra.Command{
	Use:   "results",
	Short: "list test results",
	Run:   listResultsCommand(results.NewService()),
}

func listResultsCommand(r *results.ResultsService) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		for _, result := range r.GetResultsList() {
			fmt.Println(result)
		}
	}
}
