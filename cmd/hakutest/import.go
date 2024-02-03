package main

import (
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(importCmd)
}

var importCmd = &cobra.Command{
	Use:   "import <path>",
	Short: "Import test file",
	Long:  "Import hakutest test files",
	Args:  cobra.ExactArgs(1),
	RunE:  importCommand(test.NewService()),
}

func importCommand(s *test.TestService) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return s.Import(args[0])
	}
}
