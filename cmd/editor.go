package cmd

import (
	"github.com/shelepuginivan/hakutest/internal/app/editor"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(editorCmd)
}

var editorCmd = &cobra.Command{
	Use:   "editor",
	Short: "Edit test files",
	Long:  "Edit hakutest test files",
	RunE:  editor.EditorCmd,
}
