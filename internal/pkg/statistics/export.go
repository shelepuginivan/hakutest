package statistics

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/xuri/excelize/v2"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func ExportToTable(statistics Statistics) table.Table {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow, color.Bold).SprintfFunc()

	tbl := table.New("#", "Student", "Points", "%")
	tbl.WithHeaderFormatter(headerFmt)
	tbl.WithFirstColumnFormatter(columnFmt)

	for index, entry := range statistics {
		tbl.AddRow(
			index+1,
			entry.Student,
			entry.Results.Points,
			entry.Results.Percentage,
		)
	}

	return tbl
}

func ExportToExcel(statistics Statistics, testName string) error {
	file := excelize.NewFile()

	defer file.Close()

	correctStyle, err := file.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#2cbe56"},
			Pattern: 1,
		},
	})

	if err != nil {
		return err
	}

	incorrectStyle, err := file.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#eb363e"},
			Pattern: 1,
		},
	})

	if err != nil {
		return err
	}

	borderBottom, err := file.NewStyle(&excelize.Style{
		Border: []excelize.Border{{
			Type:  "bottom",
			Style: 2,
			Color: "#000000",
		}},
	})

	if err != nil {
		return err
	}

	borderRight, err := file.NewStyle(&excelize.Style{
		Border: []excelize.Border{{
			Type:  "right",
			Style: 2,
			Color: "#000000",
		}},
	})

	if err != nil {
		return err
	}

	testResultsSheet := "Test Results"
	statisticsSheet := "Statistics"

	headers := []string{"#", "Student", "Points", "Percentage"}
	index, err := file.NewSheet(testResultsSheet)

	if err != nil {
		return err
	}

	file.SetActiveSheet(index)

	for i, header := range headers {
		column, err := excelize.ColumnNumberToName(i + 1)

		if err != nil {
			return err
		}

		cell := column + "1"

		file.SetCellStyle(testResultsSheet, cell, cell, borderBottom)
		file.SetCellValue(testResultsSheet, cell, header)
	}

	for i, entry := range statistics {
		row := i + 2

		file.SetCellValue(testResultsSheet, fmt.Sprintf("A%d", row), i+1)
		file.SetCellValue(testResultsSheet, fmt.Sprintf("B%d", row), entry.Student)
		file.SetCellValue(testResultsSheet, fmt.Sprintf("C%d", row), entry.Results.Points)
		file.SetCellValue(testResultsSheet, fmt.Sprintf("D%d", row), entry.Results.Percentage)
	}

	if _, err = file.NewSheet(statisticsSheet); err != nil {
		return err
	}

	file.SetCellValue(statisticsSheet, "A1", "#")

	for i, entry := range statistics {
		row := i + 2

		if err != nil {
			return err
		}

		studentNameCell := fmt.Sprintf("A%d", row)

		file.SetCellStyle(statisticsSheet, studentNameCell, studentNameCell, borderRight)
		file.SetCellValue(statisticsSheet, studentNameCell, entry.Student)

		for taskNumber, correct := range entry.Results.Tasks {
			taskIndex, err := strconv.Atoi(taskNumber)

			if err != nil {
				return err
			}

			column, err := excelize.ColumnNumberToName(taskIndex + 2)

			if err != nil {
				return err
			}

			valueCell := fmt.Sprintf("%s%d", column, row)
			taskNumberCell := column + "1"

			file.SetCellStyle(statisticsSheet, taskNumberCell, taskNumberCell, borderBottom)
			file.SetCellValue(statisticsSheet, taskNumberCell, taskIndex+1)

			if correct {
				file.SetCellStyle(statisticsSheet, valueCell, valueCell, correctStyle)
				file.SetCellValue(statisticsSheet, valueCell, 1)
			} else {
				file.SetCellStyle(statisticsSheet, valueCell, valueCell, incorrectStyle)
				file.SetCellValue(statisticsSheet, valueCell, 0)
			}
		}
	}

	return file.SaveAs(testName + ".xlsx")
}

func ExportToPng(statistics Statistics, testName string) error {
	p := plot.New()

	values := plotter.Values{}

	for _, entry := range statistics {
		values = append(values, float64(entry.Results.Points))
	}

	hist, err := plotter.NewHist(values, 16)

	if err != nil {
		return err
	}

	p.Add(hist)

	p.Title.Text = "Student Performance"
	p.X.Label.Text = "Points"
	p.Y.Label.Text = "Students"

	return p.Save(8*vg.Inch, 4*vg.Inch, testName+".png")
}
