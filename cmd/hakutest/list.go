package main

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "list available files",
	Aliases: []string{"ls"},
}
