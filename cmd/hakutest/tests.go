package main

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(testsCmd)
}

var testsCmd = &cobra.Command{
	Use:   "tests",
	Short: "Manage tests",
}
