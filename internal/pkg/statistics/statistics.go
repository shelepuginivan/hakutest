package statistics

import (
	"os"
	"path"

	"github.com/shelepuginivan/hakutest/internal/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/core"
	"gopkg.in/yaml.v3"
)

type Statistics = []core.TestResults

func GetStatistics(testName string) (Statistics, error) {
	stats := Statistics{}
	testResultsDir := path.Join(config.Init().General.ResultsDirectory, testName)
	entries, err := os.ReadDir(testResultsDir)

	if err != nil {
		return stats, err
	}

	for _, file := range entries {
		if file.IsDir() {
			continue
		}

		data, err := os.ReadFile(path.Join(testResultsDir, file.Name()))

		if err != nil {
			continue
		}

		entry := core.TestResults{}

		if yaml.Unmarshal(data, &entry) != nil {
			continue
		}

		stats = append(stats, entry)
	}

	return stats, nil
}
