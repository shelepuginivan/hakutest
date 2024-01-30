package statistics

import (
	"github.com/spf13/cobra"
)

func Cmd(s *StatisticsService) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		testName := args[0]

		if len(args) == 1 {
			return s.Export(testName, testName, FormatTable)
		}

		return s.Export(testName, testName, args[1])
	}
}
