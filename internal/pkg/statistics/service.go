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

func (s StatisticsService) writeResultsSheetHeaders(
	file *excelize.File,
	sheet string,
	headers []string,
) error {
	for i, header := range headers {
		column, err := excelize.ColumnNumberToName(i + 1)
		if err != nil {
			return err
		}

		cell := column + "1"

		if err := file.SetCellValue(sheet, cell, header); err != nil {
			return err
		}
	}

	return nil
}

func (s StatisticsService) writeEntryResults(
	file *excelize.File,
	entry results.TestResults,
	sheet string,
	index int,
) error {
	row := index + 2

	err := file.SetCellValue(sheet, fmt.Sprintf("A%d", row), index+1)
	if err != nil {
		return err
	}

	err = file.SetCellValue(sheet, fmt.Sprintf("B%d", row), entry.Student)
	if err != nil {
		return err
	}

	err = file.SetCellValue(sheet, fmt.Sprintf("C%d", row), entry.Results.Points)
	if err != nil {
		return err
	}

	return file.SetCellValue(sheet, fmt.Sprintf("D%d", row), entry.Results.Percentage)
}

func (s StatisticsService) writeEntryStatistics(
	file *excelize.File,
	entry results.TestResults,
	sheet string,
	index int,
	correctStyle, incorrectStyle int,
) error {
	row := index + 2
	studentNameCell := fmt.Sprintf("A%d", row)

	if err := file.SetCellValue(sheet, studentNameCell, entry.Student); err != nil {
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

		if err := file.SetCellValue(sheet, taskNumberCell, taskIndex); err != nil {
			return err
		}

		if err := file.SetCellValue(sheet, valueCell, taskResult.Answer); err != nil {
			return err
		}

		var cellStyle int

		if taskResult.Correct {
			cellStyle = correctStyle
		} else {
			cellStyle = incorrectStyle
		}

		if err := file.SetCellStyle(sheet, valueCell, valueCell, cellStyle); err != nil {
			return err
		}
	}

	return nil
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

	if _, err = file.NewSheet(statisticsSheet); err != nil {
		return err
	}

	if err := file.SetCellValue(statisticsSheet, "A1", "#"); err != nil {
		return err
	}

	if err := s.writeResultsSheetHeaders(file, testResultsSheet, headers); err != nil {
		return err
	}

	for i, entry := range stats.Entries {
		err := s.writeEntryResults(file, entry, testResultsSheet, i)
		if err != nil {
			return err
		}

		err = s.writeEntryStatistics(file, entry, statisticsSheet, i, correctStyle, incorrectStyle)
		if err != nil {
			return err
		}
	}

	defaultSheet := "Sheet1"

	if _, err = file.GetSheetIndex(defaultSheet); err == nil {
		file.DeleteSheet(defaultSheet)
	}

	return file.SaveAs(dest)
}
