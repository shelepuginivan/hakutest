package main

import "github.com/spf13/cobra"

func init() {
	resultsCmd.AddCommand(resultsListCmd)
	resultsCmd.AddCommand(resultsRemoveCmd)
}

var resultsCmd = &cobra.Command{
	Use:   "results",
	Short: "Manage test results",
}
