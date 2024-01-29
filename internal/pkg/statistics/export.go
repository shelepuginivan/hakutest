package statistics

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/xuri/excelize/v2"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const (
	FormatExcel string = "excel"
	FormatImage string = "image"
)

func (s Statistics) Export(dest string, format string) error {
	switch format {
	case FormatExcel:
		return s.ExportToExcel(dest)
	case FormatImage:
		return s.ExportToPng(dest)
	default:
		return fmt.Errorf("unknown format %s", format)
	}
}

func (s Statistics) ExportToTable() table.Table {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow, color.Bold).SprintfFunc()

	tbl := table.New("#", "Student", "Points", "%")
	tbl.WithHeaderFormatter(headerFmt)
	tbl.WithFirstColumnFormatter(columnFmt)

	for index, entry := range s.Entries {
		tbl.AddRow(
			index+1,
			entry.Student,
			entry.Results.Points,
			entry.Results.Percentage,
		)
	}

	return tbl
}

func (s Statistics) ExportToExcel(dest string) error {
	if !strings.HasSuffix(dest, ".xlsx") {
		dest += ".xlsx"
	}

	excelConfig := i18n.New().Statistics.Excel
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

	statisticsSheet := excelConfig.TestStatisticsSheet
	testResultsSheet := excelConfig.TestResultsSheet

	headers := []string{
		"#",
		excelConfig.HeaderStudent,
		excelConfig.HeaderPoints,
		excelConfig.HeaderPercentage,
	}

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

		if err := file.SetCellStyle(testResultsSheet, cell, cell, borderBottom); err != nil {
			return err
		}

		if err := file.SetCellValue(testResultsSheet, cell, header); err != nil {
			return err
		}
	}

	for i, entry := range s.Entries {
		row := i + 2

		if err := file.SetCellValue(testResultsSheet, fmt.Sprintf("A%d", row), i+1); err != nil {
			return err
		}

		if err := file.SetCellValue(testResultsSheet, fmt.Sprintf("B%d", row), entry.Student); err != nil {
			return err
		}

		if err := file.SetCellValue(testResultsSheet, fmt.Sprintf("C%d", row), entry.Results.Points); err != nil {
			return err
		}

		if err := file.SetCellValue(testResultsSheet, fmt.Sprintf("D%d", row), entry.Results.Percentage); err != nil {
			return err
		}

	}

	if _, err = file.NewSheet(statisticsSheet); err != nil {
		return err
	}

	if err := file.SetCellValue(statisticsSheet, "A1", "#"); err != nil {
		return err
	}

	for i, entry := range s.Entries {
		row := i + 2

		if err != nil {
			return err
		}

		studentNameCell := fmt.Sprintf("A%d", row)

		if err := file.SetCellStyle(statisticsSheet, studentNameCell, studentNameCell, borderRight); err != nil {
			return err
		}

		if err := file.SetCellValue(statisticsSheet, studentNameCell, entry.Student); err != nil {
			return err
		}

		for taskNumber, taskResult := range entry.Results.Tasks {
			taskIndex, err := strconv.Atoi(taskNumber)

			if err != nil {
				return err
			}

			column, err := excelize.ColumnNumberToName(taskIndex + 1)

			if err != nil {
				return err
			}

			valueCell := fmt.Sprintf("%s%d", column, row)
			taskNumberCell := column + "1"

			if err := file.SetCellStyle(statisticsSheet, taskNumberCell, taskNumberCell, borderBottom); err != nil {
				return err
			}

			if err := file.SetCellValue(statisticsSheet, taskNumberCell, taskIndex); err != nil {
				return err
			}

			if err := file.SetCellValue(statisticsSheet, valueCell, taskResult.Answer); err != nil {
				return err
			}

			var cellStyle int

			if taskResult.Correct {
				cellStyle = correctStyle
			} else {
				cellStyle = incorrectStyle
			}

			if err := file.SetCellStyle(statisticsSheet, valueCell, valueCell, cellStyle); err != nil {
				return err
			}
		}
	}

	defaultSheet := "Sheet1"

	if _, err = file.GetSheetIndex(defaultSheet); err == nil {
		file.DeleteSheet(defaultSheet)
	}

	return file.SaveAs(dest)
}

func (s Statistics) ExportToPng(dest string) error {
	if !strings.HasSuffix(dest, ".png") {
		dest += ".png"
	}

	p := plot.New()
	values := plotter.Values{}
	imageConfig := i18n.New().Statistics.Image

	for _, entry := range s.Entries {
		values = append(values, float64(entry.Results.Points))
	}

	hist, err := plotter.NewHist(values, 16)

	if err != nil {
		return err
	}

	p.Add(hist)

	p.Title.Text = imageConfig.Title
	p.X.Label.Text = imageConfig.LabelX
	p.Y.Label.Text = imageConfig.LabelY

	return p.Save(8*vg.Inch, 4*vg.Inch, dest)
}
