package parser

import (
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/shelepuginivan/hakutest/internal/config"
	"gopkg.in/yaml.v3"
)

type TestResults struct {
	Student     string    `yaml:"student"`
	SubmittedAt time.Time `yaml:"submittedAt"`
	Results     Results   `yaml:"results"`
	Test        TestInfo  `yaml:"test"`
}

type Results struct {
	Points     int             `yaml:"points"`
	Total      int             `yaml:"total"`
	Percentage int             `yaml:"percentage"`
	Tasks      map[string]bool `yaml:"tasks"`
}

type TestInfo struct {
	Title  string `yaml:"title"`
	Author string `yaml:"author"`
	Sha256 string `yaml:"sha256"`
}

func CompareAnswers(received, expected string) bool {
	return strings.TrimSpace(strings.ToLower(received)) == strings.TrimSpace(strings.ToLower(expected))
}

func (t Test) GetResults(answers map[string][]string) TestResults {
	submittedAt := time.Now()
	student := strings.Join(answers["student"], "")

	results := TestResults{
		Student:     student,
		SubmittedAt: submittedAt,
		Test: TestInfo{
			Title:  t.Title,
			Author: t.Author,
			Sha256: t.Sha256Sum(),
		},
		Results: Results{
			Points:     0,
			Total:      len(t.Tasks),
			Percentage: 0,
			Tasks:      make(map[string]bool),
		},
	}

	for i, answer := range answers {
		index, err := strconv.Atoi(i)

		if err != nil {
			continue
		}

		studentAnswer := strings.Join(answer, ",")
		correctAnswer := t.Tasks[index].Answer
		isCorrect := CompareAnswers(studentAnswer, correctAnswer)

		results.Results.Tasks[i] = isCorrect

		if isCorrect {
			results.Results.Points++
		}
	}

	results.Results.Percentage = 100 * results.Results.Points / len(t.Tasks)

	return results
}

func SaveTestResults(name string, results TestResults) error {
	testResultsDirectory := path.Join(config.Init().General.ResultsDirectory, name)
	resultsFilePath := path.Join(testResultsDirectory, results.Student+".txt")

	if _, err := os.Stat(resultsFilePath); !os.IsNotExist(err) {
		// Test was already submitted by this student
		return err
	}

	err := os.MkdirAll(testResultsDirectory, 0770)

	if err != nil && !os.IsExist(err) {
		return err
	}

	data, err := yaml.Marshal(results)

	if err != nil {
		return err
	}

	return os.WriteFile(resultsFilePath, data, 0666)
}
