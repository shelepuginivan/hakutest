package statistics

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func ExportToExcel(statistics Statistics, testName string) error {
	file := excelize.NewFile()

	defer file.Close()

	sheetName := "Test Results"
	headers := []string{"#", "Student", "Points", "Percentage"}
	index, err := file.NewSheet(sheetName)

	if err != nil {
		return err
	}

	file.SetActiveSheet(index)

	for i, header := range headers {
		column, err := excelize.ColumnNumberToName(i + 1)

		if err != nil {
			return err
		}

		file.SetCellValue(sheetName, column+"1", header)
	}

	for i, entry := range statistics {
		row := i + 2

		file.SetCellValue(sheetName, fmt.Sprintf("A%d", row), i+1)
		file.SetCellValue(sheetName, fmt.Sprintf("B%d", row), entry.Student)
		file.SetCellValue(sheetName, fmt.Sprintf("C%d", row), entry.Results.Points)
		file.SetCellValue(sheetName, fmt.Sprintf("D%d", row), entry.Results.Percentage)
	}

	return file.SaveAs(testName + ".xlsx")
}
