package statistics

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/shelepuginivan/hakutest/internal/pkg/application"
	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/xuri/excelize/v2"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// StatisticsService is a struct that provides methods for manipulating Statistics structures.
type StatisticsService struct {
	app *application.App
	r   *results.ResultsService
}

// NewService returns a StatisticsService instance.
func NewService(app *application.App, r *results.ResultsService) *StatisticsService {
	return &StatisticsService{
		app: app,
		r:   r,
	}
}

// Export retrieves statistics of the test and exports in to a specified format.
func (s StatisticsService) Export(testName, dest, format string) error {
	res, err := s.r.GetResultsOfTest(testName)

	if err != nil {
		return err
	}

	stats := New(res)

	switch format {
	case FormatExcel:
		return s.ExportToExcel(stats, dest)
	case FormatImage:
		return s.ExportToPng(stats, dest)
	case FormatTable:
		s.ExportToTable(stats).Print()
		return nil
	default:
		return fmt.Errorf("unknown format %s", format)
	}
}

func (s StatisticsService) ExportToTable(stats *Statistics) table.Table {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow, color.Bold).SprintfFunc()

	tbl := table.New("#", "Student", "Points", "%")
	tbl.WithHeaderFormatter(headerFmt)
	tbl.WithFirstColumnFormatter(columnFmt)

	for index, entry := range stats.Entries {
		tbl.AddRow(
			index+1,
			entry.Student,
			entry.Results.Points,
			entry.Results.Percentage,
		)
	}

	return tbl
}

func (s StatisticsService) ExportToExcel(stats *Statistics, dest string) error {
	if !strings.HasSuffix(dest, ".xlsx") {
		dest += ".xlsx"
	}

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

	statisticsSheet := s.app.I18n.Statistics.Excel.TestStatisticsSheet
	testResultsSheet := s.app.I18n.Statistics.Excel.TestResultsSheet

	headers := []string{
		"#",
		s.app.I18n.Statistics.Excel.HeaderStudent,
		s.app.I18n.Statistics.Excel.HeaderPoints,
		s.app.I18n.Statistics.Excel.HeaderPercentage,
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

	for i, entry := range stats.Entries {
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

	for i, entry := range stats.Entries {
		row := i + 2

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

func (s StatisticsService) ExportToPng(stats *Statistics, dest string) error {
	if !strings.HasSuffix(dest, ".png") {
		dest += ".png"
	}

	p := plot.New()
	values := plotter.Values{}

	for _, entry := range stats.Entries {
		values = append(values, float64(entry.Results.Points))
	}

	hist, err := plotter.NewHist(values, 16)

	if err != nil {
		return err
	}

	p.Add(hist)

	p.Title.Text = s.app.I18n.Statistics.Image.Title
	p.X.Label.Text = s.app.I18n.Statistics.Image.LabelX
	p.Y.Label.Text = s.app.I18n.Statistics.Image.LabelY

	return p.Save(8*vg.Inch, 4*vg.Inch, dest)
}
