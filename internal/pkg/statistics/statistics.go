package statistics

import (
	"os"
	"path/filepath"

	"github.com/shelepuginivan/hakutest/internal/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/core"
	"gopkg.in/yaml.v3"
)

type Statistics struct {
	Entries []core.TestResults
}

func GetStatistics(testName string) (Statistics, error) {
	stats := Statistics{}
	testResultsDir := filepath.Join(config.Init().General.ResultsDirectory, testName)
	entries, err := os.ReadDir(testResultsDir)

	if err != nil {
		return stats, err
	}

	for _, file := range entries {
		if file.IsDir() {
			continue
		}

		data, err := os.ReadFile(filepath.Join(testResultsDir, file.Name()))

		if err != nil {
			continue
		}

		entry := core.TestResults{}

		if yaml.Unmarshal(data, &entry) != nil {
			continue
		}

		stats.Entries = append(stats.Entries, entry)
	}

	return stats, nil
}
