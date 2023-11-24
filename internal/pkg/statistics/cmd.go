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
