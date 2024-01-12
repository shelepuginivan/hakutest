package core

import (
	"os"
	"path/filepath"
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

type TaskResult struct {
	Answer  string `yaml:"answer"`
	Correct bool   `yaml:"correct"`
}

type Results struct {
	Points     int                   `yaml:"points"`
	Total      int                   `yaml:"total"`
	Percentage int                   `yaml:"percentage"`
	Tasks      map[string]TaskResult `yaml:"tasks"`
}

type TestInfo struct {
	Title  string `yaml:"title"`
	Author string `yaml:"author"`
	Sha256 string `yaml:"sha256"`
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
		isCorrect := CompareAnswers(studentAnswer, correctAnswer)

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

func (r TestResults) Save(name string) error {
	testName := strings.TrimSuffix(name, ".json")
	testResultsDirectory := filepath.Join(config.New().General.ResultsDirectory, testName)
	resultsFilePath := filepath.Join(testResultsDirectory, r.Student+".txt")

	if _, err := os.Stat(resultsFilePath); !os.IsNotExist(err) {
		// Test was already submitted by this student
		return err
	}

	err := os.MkdirAll(testResultsDirectory, 0770)

	if err != nil && !os.IsExist(err) {
		return err
	}

	data, err := yaml.Marshal(r)

	if err != nil {
		return err
	}

	return os.WriteFile(resultsFilePath, data, 0666)
}
