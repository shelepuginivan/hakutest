package main

import (
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
	"github.com/spf13/cobra"
)

func importCommand(cmd *cobra.Command, args []string) error {
	return test.NewService(app).Import(args[0])
}

var importCmd = &cobra.Command{
	Use:   "import <path>",
	Short: "Import test file",
	Long:  "Import hakutest test files",
	Args:  cobra.ExactArgs(1),
	RunE:  importCommand,
}
