package results

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	defer setup()()
	m.Run()
}

func setup() func() {
	generalConfig := config.New().General
	testName := "__mock__"

	mockTest := test.Test{
		Title:  "Mock test",
		Author: "John Doe",
		Tasks: []test.Task{
			{Answer: "A"},
			{Answer: "B"},
			{Answer: "C"},
		},
	}

	err := test.NewService().SaveToTestsDirectory(mockTest, testName)
	if err != nil {
		panic(err)
	}

	return func() {
		testPath := filepath.Join(generalConfig.TestsDirectory, testName+".json")
		resultsPath := filepath.Join(generalConfig.ResultsDirectory, testName)

		err := os.Remove(testPath)
		if err != nil {
			panic(err)
		}

		err = os.RemoveAll(resultsPath)
		if err != nil {
			panic(err)
		}
	}
}

func TestResultsService_CompareAnswers(t *testing.T) {
	service := NewService()
	cases := []struct {
		a        string
		b        string
		expected bool
	}{
		{a: "a", b: "a", expected: true},
		{a: "A", b: "a", expected: true},
		{a: "a", b: "A", expected: true},
		{a: "aaaa", b: "aAaA", expected: true},
		{a: "abABabAB", b: "aBaBaBaB", expected: true},
		{a: "baobab", b: "baObAB", expected: true},
		{a: "a ", b: "a", expected: true},
		{a: " a", b: "a", expected: true},
		{a: " a ", b: "a", expected: true},
		{a: "a", b: "a ", expected: true},
		{a: "a", b: " a", expected: true},
		{a: "a", b: " a ", expected: true},
		{a: "a             ", b: "             a", expected: true},
		{a: "different", b: "answer", expected: false},
		{a: "racecar", b: "ricecar", expected: false},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, service.CompareAnswers(c.a, c.b))
	}
}

func TestResultsService_CheckAnswers(t *testing.T) {
	s := ResultsService{}

	mockTest := test.Test{
		Title:  "Mock Test",
		Author: "John Doe",
		Tasks: []test.Task{
			{Answer: "A"},
			{Answer: "B"},
			{Answer: "C"},
		},
	}

	mockAnswers := map[string][]string{
		"student": {"ABC"},
		"0":       {"A"},
		"1":       {"B"},
		"2":       {"C"},
	}

	results := s.CheckAnswers(mockTest, mockAnswers)

	expectedResultsInfo := TestResults{
		Student:     "ABC",
		SubmittedAt: results.SubmittedAt,
		Test: TestInfo{
			Title:  mockTest.Title,
			Author: mockTest.Author,
			Sha256: mockTest.Sha256Sum(),
		},
		Results: Results{
			Total:      3,
			Points:     3,
			Percentage: 100,
			Tasks: map[string]TaskResult{
				"1": {Answer: "A", Correct: true},
				"2": {Answer: "B", Correct: true},
				"3": {Answer: "C", Correct: true},
			},
		},
	}

	assert.Equal(t, expectedResultsInfo, results)
	assert.Less(t, expectedResultsInfo.SubmittedAt, time.Now())
}

func TestResultsService_Save(t *testing.T) {
	mockResults := TestResults{
		Student:     "ABC",
		SubmittedAt: time.Now(),
		Test: TestInfo{
			Title:  "Mock test",
			Author: "John Doe",
			Sha256: "sha256",
		},
		Results: Results{
			Total:      3,
			Points:     3,
			Percentage: 100,
			Tasks: map[string]TaskResult{
				"1": {Answer: "A", Correct: true},
				"2": {Answer: "B", Correct: true},
				"3": {Answer: "C", Correct: true},
			},
		},
	}

	assert.NoError(t, NewService().Save(mockResults, "__mock__"))
}

func TestResultsService_GetResultsOfTest(t *testing.T) {
	results, err := NewService().GetResultsOfTest("__mock__")

	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.IsType(t, TestResults{}, results[0])
}
