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
	Use:       "statistics <test> [format]",
	Short:     "Test results statistics",
	Long:      "Export test results statistics",
	Args:      cobra.RangeArgs(1, 2),
	ValidArgs: []string{statistics.FormatExcel, statistics.FormatImage, statistics.FormatTable},
	RunE:      statisticsCommand,
	Aliases:   []string{"stats"},
}

func statisticsCommand(cmd *cobra.Command, args []string) error {
	testName := args[0]
	r := results.NewService(app)
	s := statistics.NewService(app, r)

	if len(args) == 1 {
		return s.Export(testName, testName, statistics.FormatTable)
	}

	return s.Export(testName, testName, args[1])
}
