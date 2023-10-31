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

		fmt.Println(correctAnswer, studentAnswer)

		if CompareAnswers(studentAnswer, correctAnswer) {
			points++
		}
	}

	percentage := 100 * points / len(test.Tasks)

	results = TestResults{
		Student:     student,
		SubmittedAt: submittedAt,
		Points:      points,
		Percentage:  percentage,
	}

	return results, nil
}
