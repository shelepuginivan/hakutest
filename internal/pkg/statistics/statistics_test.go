package statistics

import (
	"os"
	"testing"
	"time"

	parser "github.com/shelepuginivan/hakutest/internal/pkg/test_parser"
)

func TestExportToPng(t *testing.T) {
	statistics := Statistics{
		parser.TestResults{
			Student:     "John Doe",
			SubmittedAt: time.Now(),
			Results: parser.Results{
				Points:     80,
				Total:      100,
				Percentage: 80,
				Tasks: map[string]bool{
					"0": false,
					"1": true,
					"2": false,
				},
			},
			Test: parser.TestInfo{
				Title:  "Unit Test 1",
				Author: "Jane Smith",
				Sha256: "abcdef123456",
			},
		},
	}

	outputFile := "test_output"

	if ExportToPng(statistics, outputFile) != nil {
		t.Fail()
	}

	defer func() {
		err := os.Remove(outputFile + ".png")
		if err != nil {
			t.Logf("Failed to delete file: %s", err)
		}
	}()
}

func TestExportToExcel(t *testing.T) {
	statistics := Statistics{
		parser.TestResults{
			Student:     "Alex",
			SubmittedAt: time.Now(),
			Results: parser.Results{
				Points:     80,
				Total:      100,
				Percentage: 80,
				Tasks: map[string]bool{
					"0": true,
					"1": false,
					"2": true,
					"3": true,
					"4": true,
					"5": false,
				},
			},
			Test: parser.TestInfo{
				Title:  "Unit Test 2",
				Author: "John Doe",
				Sha256: "53d4843cec3163136498a8e570dcde153046ec0009e57bde714649f95aebe7f7",
			},
		},
	}

	outputFile := "test_output"

	if err := ExportToExcel(statistics, outputFile); err != nil {
		t.Fail()
	}

	defer func() {
		err := os.Remove(outputFile + ".xlsx")
		if err != nil {
			t.Logf("Failed to delete file: %s", err)
		}
	}()
}
