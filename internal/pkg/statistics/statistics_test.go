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
		Entries: []core.TestResults{
			{
				Student:     "John Doe",
				SubmittedAt: time.Now(),
				Results: core.Results{
					Points:     80,
					Total:      100,
					Percentage: 80,
					Tasks: map[string]core.TaskResult{
						"0": {Answer: "4", Correct: false},
						"1": {Answer: "1,2,5", Correct: true},
						"2": {Answer: "some", Correct: true},
					},
				},
				Test: core.TestInfo{
					Title:  "Unit Test 1",
					Author: "Jane Smith",
					Sha256: "abcdef123456",
				},
			},
		},
	}

	outputFile := "test_output"

	assert.Nil(t, statistics.ExportToPng(outputFile))

	defer func() {
		err := os.Remove(outputFile + ".png")
		if err != nil {
			t.Logf("Failed to delete file: %s", err)
		}
	}()
}

func TestExportToExcel(t *testing.T) {
	statistics := Statistics{
		Entries: []core.TestResults{
			{
				Student:     "Alex",
				SubmittedAt: time.Now(),
				Results: core.Results{
					Points:     80,
					Total:      100,
					Percentage: 80,
					Tasks: map[string]core.TaskResult{
						"0": {Answer: "1,2,4", Correct: true},
						"1": {Answer: "answer", Correct: false},
						"2": {Answer: "3", Correct: true},
						"3": {Answer: "2,4", Correct: true},
						"4": {Answer: "some", Correct: true},
						"5": {Answer: "another", Correct: false},
					},
				},
				Test: core.TestInfo{
					Title:  "Unit Test 2",
					Author: "John Doe",
					Sha256: "53d4843cec3163136498a8e570dcde153046ec0009e57bde714649f95aebe7f7",
				},
			},
		},
	}

	outputFile := "test_output"

	assert.Nil(t, statistics.ExportToExcel(outputFile))

	defer func() {
		err := os.Remove(outputFile + ".xlsx")
		if err != nil {
			t.Logf("Failed to delete file: %s", err)
		}
	}()
}
