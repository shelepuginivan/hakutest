package statistics

import (
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

func Cmd(cmd *cobra.Command, args []string) error {
	stats, err := GetStatistics(args[0])

	if err != nil {
		return err
	}

	if len(args) == 1 {
		tbl := table.New("#", "Student", "Points", "%")

		for index, entry := range stats {
			tbl.AddRow(
				index+1,
				entry.Student,
				entry.Results.Points,
				entry.Results.Percentage,
			)
		}

		tbl.Print()

		return nil
	}

	switch args[1] {
	case "excel":
		return ExportToExcel(stats, args[0])
	}

	return nil
}
