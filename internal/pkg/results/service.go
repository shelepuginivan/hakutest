package results

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
	"gopkg.in/yaml.v3"
)

// ResultsService is a struct that provides methods for manipulating Result structures.
type ResultsService struct{}

// NewService returns a ResultsService instance.
func NewService() *ResultsService {
	return &ResultsService{}
}

// CompareAnswers reports whether answers match.
// It is case-insensitive and ignores leading and trailing spaces.
func (s ResultsService) CompareAnswers(received, expected string) bool {
	return strings.TrimSpace(strings.ToLower(received)) == strings.TrimSpace(strings.ToLower(expected))
}

// CheckAnswers returns the results of the test.
func (s ResultsService) CheckAnswers(t test.Test, answers map[string][]string) TestResults {
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
			Tasks:      make(map[string]TaskResult),
		},
	}

	for i, answer := range answers {
		index, err := strconv.Atoi(i)

		if err != nil {
			continue
		}

		studentAnswer := strings.Join(answer, ",")
		correctAnswer := t.Tasks[index].Answer
		isCorrect := s.CompareAnswers(studentAnswer, correctAnswer)

		results.Results.Tasks[strconv.Itoa(index+1)] = TaskResult{
			Answer:  studentAnswer,
			Correct: isCorrect,
		}

		if isCorrect {
			results.Results.Points++
		}
	}

	results.Results.Percentage = 100 * results.Results.Points / len(t.Tasks)

	return results
}

// GetResultsList retrieves a list of results names from the results directory specified in the configuration.
func (s ResultsService) GetResultsList() []string {
	resultsDirectory := config.New().General.ResultsDirectory
	resultsList := []string{}

	entries, err := os.ReadDir(resultsDirectory)

	if err != nil {
		return resultsList
	}

	for _, file := range entries {
		resultsName := file.Name()

		if file.IsDir() {
			resultsList = append(resultsList, resultsName)
		}
	}

	return resultsList
}

// GetResultsOfTest retrieves all results of the test from the results directory specified in the configuration.
// The name is the filename of the test.
func (s ResultsService) GetResultsOfTest(name string) ([]TestResults, error) {
	results := []TestResults{}
	testResultsDir := filepath.Join(config.New().General.ResultsDirectory, name)
	entries, err := os.ReadDir(testResultsDir)

	if err != nil {
		return nil, err
	}

	for _, file := range entries {
		if file.IsDir() {
			continue
		}

		data, err := os.ReadFile(filepath.Join(testResultsDir, file.Name()))

		if err != nil {
			continue
		}

		entry := TestResults{}

		if yaml.Unmarshal(data, &entry) != nil {
			continue
		}

		results = append(results, entry)
	}

	return results, nil
}

// Save saves test results in the results directory specified in the configuration.
// The results are saved in a subdirectory name.
func (s ResultsService) Save(r TestResults, name string) error {
	generalConfig := config.New().General
	testName := strings.TrimSuffix(name, ".json")
	testResultsDirectory := filepath.Join(generalConfig.ResultsDirectory, testName)
	resultsFilePath := filepath.Join(testResultsDirectory, r.Student+".txt")

	_, err := os.Stat(resultsFilePath)
	if !os.IsNotExist(err) && !generalConfig.OverwriteResults {
		// Test was already submitted by this student
		return err
	}

	err = os.MkdirAll(testResultsDirectory, os.ModeDir|os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return err
	}

	data, err := yaml.Marshal(r)
	if err != nil {
		return err
	}

	return os.WriteFile(resultsFilePath, data, 0666)
}
