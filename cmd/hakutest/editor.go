package main

import (
	"github.com/shelepuginivan/hakutest/internal/app/editor"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(editorCmd)
}

var editorCmd = &cobra.Command{
	Use:     "editor [filename]",
	Short:   "Edit test files",
	Long:    "Edit hakutest test files",
	Args:    cobra.RangeArgs(0, 1),
	RunE:    editor.Cmd(test.NewService()),
	Aliases: []string{"e"},
}
