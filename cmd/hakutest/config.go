package main

import (
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/spf13/cobra"
)

func configCommand(c *config.ConfigService) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return c.PrintConfig()
		}

		if len(args) == 1 {
			return c.PrintField(args[0])
		}

		return c.SetField(args[0], args[1])
	}
}

var configCmd = &cobra.Command{
	Use:     "config [field] [value]",
	Short:   "Manage the configuration settings",
	Long:    "Manage hakutest configuration settings",
	Args:    cobra.RangeArgs(0, 2),
	RunE:    configCommand(config.NewService()),
	Aliases: []string{"cfg"},
}
