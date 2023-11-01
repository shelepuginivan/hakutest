package statistics

import (
	"github.com/spf13/cobra"
)

func Cmd(cmd *cobra.Command, args []string) error {
	stats, err := GetStatistics(args[0])

	if err != nil {
		return err
	}

	if len(args) == 1 {
		ExportToTable(stats).Print()
		return nil
	}

	switch args[1] {
	case "excel":
		return ExportToExcel(stats, args[0])
	default:
		ExportToTable(stats).Print()
		return nil
	}
}
