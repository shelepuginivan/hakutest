package statistics

import (
	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/spf13/cobra"
)

func Cmd(s results.ResultsService) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		results, err := s.GetResultsOfTest(args[0])

		if err != nil {
			return err
		}

		stats := New(results)

		if len(args) == 1 {
			stats.ExportToTable().Print()
			return nil
		}

		switch args[1] {
		case "excel":
			return stats.ExportToExcel(args[0])
		case "image":
			return stats.ExportToPng(args[0])
		default:
			stats.ExportToTable().Print()
			return nil
		}
	}
}
