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

func GetTestResults(name string, answers map[string][]string) (TestResults, error) {
	submittedAt := time.Now()
	student := strings.Join(answers["student"], "")
	test, err := ParseTest(name)

	if err != nil {
		return TestResults{}, err
	}

	checksum, err := GetTestCheckSum(name)

	if err != nil {
		return TestResults{}, err
	}

	results := TestResults{
		Student:     student,
		SubmittedAt: submittedAt,
		Test: TestInfo{
			Title:  test.Title,
			Author: test.Author,
			Sha256: checksum,
		},
		Results: Results{
			Points:     0,
			Total:      len(test.Tasks),
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
		correctAnswer := test.Tasks[index].Answer
		isCorrect := CompareAnswers(studentAnswer, correctAnswer)

		results.Results.Tasks[i] = isCorrect

		if isCorrect {
			results.Results.Points++
		}
	}

	results.Results.Percentage = 100 * results.Results.Points / len(test.Tasks)

	return results, nil
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
