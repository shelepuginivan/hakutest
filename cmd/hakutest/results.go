package main

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(resultsCmd)
}

var resultsCmd = &cobra.Command{
	Use:   "results",
	Short: "Manage test results",
}
