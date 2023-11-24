package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareAnswer(t *testing.T) {
	cases := []struct {
		a        string
		b        string
		expected bool
	}{
		{a: "Answer", b: "Answer", expected: true},
		{a: "Answer", b: "answer", expected: true},
		{a: "ANSWER", b: "Answer", expected: true},
		{a: "1 2 3 4 5 ", b: "1 2 3 4 5", expected: true},
		{a: "   a   ", b: "A", expected: true},
		{a: "", b: "", expected: true},
		{a: "   ", b: "", expected: true},
		{a: "B", b: "b", expected: true},
		{a: "1", b: "2", expected: false},
		{a: "   1   ", b: "2", expected: false},
		{a: "   Some   ", b: "  Another  ", expected: false},
	}

	for _, c := range cases {
		assert.Equal(t, CompareAnswers(c.a, c.b), c.expected)
	}
}

func TestGetResults(t *testing.T) {
	test := Test{
		Tasks: []Task{
			{
				Type:    "multiple",
				Options: []string{"8", "9", "10"},
				Answer:  "1,2",
			},
			{
				Type:    "single",
				Options: []string{"b", "c", "d", "e"},
				Answer:  "2",
			},
			{
				Type:   "open",
				Answer: "some",
			},
			{
				Type:   "open",
				Answer: "another",
			},
		},
	}

	testInfo := TestInfo{Title: test.Title, Author: test.Author, Sha256: test.Sha256Sum()}

	cases := []struct {
		answer  map[string][]string
		results TestResults
	}{
		{
			answer: map[string][]string{
				"student": {"John Doe"},
				"0":       {"1", "2"},
				"1":       {"2"},
				"2":       {"some"},
				"3":       {"another"},
			},
			results: TestResults{
				Student: "John Doe",
				Results: Results{
					Points:     4,
					Total:      4,
					Percentage: 100,
					Tasks: map[string]bool{
						"1": true,
						"2": true,
						"3": true,
						"4": true,
					},
				},
				Test: testInfo,
			},
		},
	}

	for _, c := range cases {
		received := test.GetResults(c.answer)

		assert.Equal(t, received.Student, c.results.Student)
		assert.EqualValues(t, received.Results, c.results.Results)
		assert.EqualValues(t, received.Test, c.results.Test)
	}
}
