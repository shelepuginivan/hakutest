package cmd

import (
	"github.com/shelepuginivan/hakutest/internal/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use:   "config [field] [value]",
	Short: "Manage the configuration settings",
	Long:  "Manage hakutest configuration settings",
	Args:  cobra.RangeArgs(0, 2),
	Run:   config.Cmd,
}
