package parser

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type TestResults struct {
	Student     string
	SubmittedAt time.Time
	Points      int
	Percentage  int
}

func CompareAnswers(received, expected string) bool {
	return strings.TrimSpace(strings.ToLower(received)) == strings.TrimSpace(strings.ToLower(expected))
}

func GetTestResults(name string, answers map[string][]string) (TestResults, error) {
	results := TestResults{}
	test, err := ParseTest(name)
	submittedAt := time.Now()
	points := 0

	if err != nil {
		return results, err
	}

	student := strings.Join(answers["student"], "")

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
