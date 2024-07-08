package statistics

import (
	"fmt"
	"io"

	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/shelepuginivan/hakutest/pkg/test"
	"github.com/xuri/excelize/v2"
)

const (
	// XLSX statistics export format.
	FormatXLSX = "xlsx"

	// XLSX file MIME type.
	MimeXLSX = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
)

const (
	defaultSheet = "Sheet1"
	sheet        = "Hakutest"
)

const (
	cellBgCorrect   = "#A4D2AA"
	cellFgCorrect   = "#2C7135"
	cellBgDetailed  = "#DDD645"
	cellFgDetailed  = "#38201D"
	cellBgIncorrect = "#D68282"
	cellFgIncorrect = "#870E0F"
)

// WriteXLSX writes statistics in XLSX format to w.
func (s *Statistics) WriteXLSX(w io.Writer) error {
	f := excelize.NewFile()

	correct, incorrect, detailed, err := registerStyles(f)
	if err != nil {
		return err
	}

	sheetIndex, err := f.NewSheet(sheet)
	if err != nil {
		return err
	}
	f.SetActiveSheet(sheetIndex)

	if err = f.DeleteSheet(defaultSheet); err != nil {
		return err
	}

	if err = writeHeaderRow(f, s.Total); err != nil {
		return err
	}

	for i, r := range s.Results {
		row := i + 2 // Rows are 1-indexed, 1st row is the header row.

		if f.SetCellStr(sheet, fmt.Sprintf("A%d", row), r.Student) != nil {
			continue
		}

		if f.SetCellInt(sheet, fmt.Sprintf("B%d", row), r.Points) != nil {
			continue
		}

		if f.SetCellInt(sheet, fmt.Sprintf("C%d", row), r.Percentage) != nil {
			continue
		}

		for idx, a := range r.Answers {
			cell, err := excelize.CoordinatesToCellName(idx+4, row)
			if err != nil {
				continue
			}

			style := incorrect

			if a.Correct {
				style = correct
			}

			if a.Type == test.TaskDetailed {
				style = detailed
			}

			if f.SetCellStr(sheet, cell, a.Value) != nil {
				continue
			}

			f.SetCellStyle(sheet, cell, cell, style)
		}
	}

	return f.Write(w)
}

// registerStyles adds new styles for correct, incorrect, and detailed answers.
// It returns indices of registered styles in the above order and error if any.
func registerStyles(f *excelize.File) (correct, incorrect, detailed int, err error) {
	correct, err = f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Color:   []string{cellBgCorrect},
			Pattern: 1,
			Type:    "pattern",
		},
		Font: &excelize.Font{
			Bold:  true,
			Color: cellFgCorrect,
		},
	})

	if err != nil {
		return correct, incorrect, detailed, err
	}

	incorrect, err = f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Color:   []string{cellBgIncorrect},
			Pattern: 1,
			Type:    "pattern",
		},
		Font: &excelize.Font{
			Bold:  true,
			Color: cellFgIncorrect,
		},
	})

	if err != nil {
		return correct, incorrect, detailed, err
	}

	detailed, err = f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			// Dotted border.
			{Color: cellFgDetailed, Style: 4, Type: "left"},
			{Color: cellFgDetailed, Style: 4, Type: "top"},
			{Color: cellFgDetailed, Style: 4, Type: "right"},
			{Color: cellFgDetailed, Style: 4, Type: "bottom"},
		},
		Fill: excelize.Fill{
			Color:   []string{cellBgDetailed},
			Pattern: 1,
			Type:    "pattern",
		},
		Font: &excelize.Font{
			Color: cellFgDetailed,
		},
	})

	return correct, incorrect, detailed, err
}

// writeHeaderRow writes header row of the table.
func writeHeaderRow(f *excelize.File, totalTasks int) (err error) {
	err = f.SetCellStr(sheet, "A1", i18n.Get("statistics.view.student"))
	if err != nil {
		return err
	}

	err = f.SetCellStr(sheet, "B1", i18n.Get("statistics.view.points"))
	if err != nil {
		return err
	}

	err = f.SetCellStr(sheet, "C1", "%")
	if err != nil {
		return err
	}

	for i := range totalTasks {
		cell, err := excelize.CoordinatesToCellName(i+4, 1)
		if err != nil {
			return err
		}

		f.SetCellInt(sheet, cell, i+1)
	}

	return nil
}
