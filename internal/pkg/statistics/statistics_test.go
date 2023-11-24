package statistics

import (
	"os"
	"testing"
	"time"

	"github.com/shelepuginivan/hakutest/internal/pkg/core"
	"github.com/stretchr/testify/assert"
)

func TestExportToPng(t *testing.T) {
	statistics := Statistics{
		core.TestResults{
			Student:     "John Doe",
			SubmittedAt: time.Now(),
			Results: core.Results{
				Points:     80,
				Total:      100,
				Percentage: 80,
				Tasks: map[string]bool{
					"0": false,
					"1": true,
					"2": false,
				},
			},
			Test: core.TestInfo{
				Title:  "Unit Test 1",
				Author: "Jane Smith",
				Sha256: "abcdef123456",
			},
		},
	}

	outputFile := "test_output"

	assert.Nil(t, ExportToPng(statistics, outputFile))

	defer func() {
		err := os.Remove(outputFile + ".png")
		if err != nil {
			t.Logf("Failed to delete file: %s", err)
		}
	}()
}

func TestExportToExcel(t *testing.T) {
	statistics := Statistics{
		core.TestResults{
			Student:     "Alex",
			SubmittedAt: time.Now(),
			Results: core.Results{
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
			Test: core.TestInfo{
				Title:  "Unit Test 2",
				Author: "John Doe",
				Sha256: "53d4843cec3163136498a8e570dcde153046ec0009e57bde714649f95aebe7f7",
			},
		},
	}

	outputFile := "test_output"

	assert.Nil(t, ExportToExcel(statistics, outputFile))

	defer func() {
		err := os.Remove(outputFile + ".xlsx")
		if err != nil {
			t.Logf("Failed to delete file: %s", err)
		}
	}()
}
