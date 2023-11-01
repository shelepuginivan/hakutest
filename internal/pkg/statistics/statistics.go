package statistics

import (
	"os"
	"path"

	"github.com/shelepuginivan/hakutest/internal/config"
	parser "github.com/shelepuginivan/hakutest/internal/pkg/test_parser"
	"gopkg.in/yaml.v3"
)

func GetStatistics(testName string) ([]parser.TestResults, error) {
	stats := []parser.TestResults{}
	testResultsDir := path.Join(config.Init().ResultsDirectory, testName)
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

		entry := parser.TestResults{}

		if yaml.Unmarshal(data, &entry) != nil {
			continue
		}

		stats = append(stats, entry)
	}

	return stats, nil
}
