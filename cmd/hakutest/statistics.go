package main

import (
	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/shelepuginivan/hakutest/internal/pkg/statistics"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statisticsCmd)
}

var statisticsCmd = &cobra.Command{
	Use:   "statistics <test> [format]",
	Short: "Test results statistics",
	Long:  "Export test results statistics",
	Args:  cobra.RangeArgs(1, 2),
	RunE:  statistics.Cmd(statistics.NewService(results.NewService())),
}
