package main

import (
	"github.com/spf13/cobra"
)

func init() {
	testsCmd.AddCommand(importCmd)
	testsCmd.AddCommand(testsListCmd)
	testsCmd.AddCommand(testsRemoveCmd)
}

var testsCmd = &cobra.Command{
	Use:   "tests",
	Short: "Manage tests",
}
