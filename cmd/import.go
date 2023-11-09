package cmd

import (
	parser "github.com/shelepuginivan/hakutest/internal/pkg/test_parser"
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
	RunE:  Cmd,
}

func Cmd(cmd *cobra.Command, args []string) error {
	return parser.Import(args[0])
}