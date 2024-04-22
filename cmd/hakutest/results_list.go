package main

import (
	"fmt"

	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/spf13/cobra"
)

func resultsListCommand(cmd *cobra.Command, args []string) {
	for _, result := range results.NewService(app).GetResultsList() {
		fmt.Println(result)
	}
}

var resultsListCmd = &cobra.Command{
	Use:     "list",
	Short:   "list available results",
	Run:     resultsListCommand,
	Aliases: []string{"ls"},
}
