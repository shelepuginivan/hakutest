// Package cli provides CLI interface for the `hakuctl` binary.
package cli

import (
	"github.com/fatih/color"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/pkg/result"
	"github.com/shelepuginivan/hakutest/pkg/test"
	"github.com/spf13/cobra"
)

func init() {
	cfg := config.New()
	result.Init(cfg.Result)
	test.Init(cfg.Test)

	rootCmd.PersistentFlags().BoolVar(&color.NoColor, "no-color", false, "Disable color output")
}

var rootCmd = &cobra.Command{
	Use:   "hakuctl",
	Short: "hakuctl - control utility for Hakutest",
	Long:  "A command line interface to control Hakutest and manage its data",
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
