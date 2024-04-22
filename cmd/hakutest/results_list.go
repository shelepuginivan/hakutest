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
	Run:     resultsListCommand,
	Aliases: []string{"ls"},
}

func resultsListCommand(cmd *cobra.Command, args []string) {
	for _, result := range results.NewService(app).GetResultsList() {
		fmt.Println(result)
	}
}
