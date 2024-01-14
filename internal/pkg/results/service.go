package results

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/shelepuginivan/hakutest/internal/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
	"gopkg.in/yaml.v3"
)

type ResultsService struct{}

func NewService() ResultsService {
	return ResultsService{}
}

func (s ResultsService) CompareAnswers(received, expected string) bool {
	return strings.TrimSpace(strings.ToLower(received)) == strings.TrimSpace(strings.ToLower(expected))
}

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

func (s ResultsService) Save(r TestResults, name string) error {
	testName := strings.TrimSuffix(name, ".json")
	testResultsDirectory := filepath.Join(config.New().General.ResultsDirectory, testName)
	resultsFilePath := filepath.Join(testResultsDirectory, r.Student+".txt")

	if _, err := os.Stat(resultsFilePath); !os.IsNotExist(err) {
		// Test was already submitted by this student
		return err
	}

	err := os.MkdirAll(testResultsDirectory, os.ModeDir|os.ModePerm)

	if err != nil && !os.IsExist(err) {
		return err
	}

	data, err := yaml.Marshal(r)

	if err != nil {
		return err
	}

	return os.WriteFile(resultsFilePath, data, 0666)
}
