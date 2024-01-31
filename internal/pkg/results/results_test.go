package results

import (
	"testing"
	"time"

	"github.com/shelepuginivan/hakutest/internal/pkg/test"
	"github.com/stretchr/testify/assert"
)

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

	expectedStudent := "ABC"
	expectedTestInfo := TestInfo{
		Title:  "Mock Test",
		Author: "John Doe",
		Sha256: mockTest.Sha256Sum(),
	}
	expectedTaskResults := map[string]TaskResult{
		"1": {Answer: "A", Correct: true},
		"2": {Answer: "B", Correct: true},
		"3": {Answer: "C", Correct: true},
	}
	expectedResults := Results{
		Points:     3,
		Total:      3,
		Percentage: 100,
		Tasks:      expectedTaskResults,
	}
	expectedResultsInfo := TestResults{
		Student:     expectedStudent,
		SubmittedAt: results.SubmittedAt,
		Test:        expectedTestInfo,
		Results:     expectedResults,
	}

	assert.Equal(t, expectedResultsInfo, results)
	assert.Less(t, expectedResultsInfo.SubmittedAt, time.Now())
}
