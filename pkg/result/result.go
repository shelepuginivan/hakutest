// Package result provides results management for the app.
package result

import (
	"time"

	"github.com/shelepuginivan/hakutest/pkg/test"
)

// Answer represents answer given by the student.
type Answer struct {
	Type    string `json:"type"`    // Type of the task.
	Value   string `json:"value"`   // The answer.
	Correct bool   `json:"correct"` // Whether answer is correct.
}

// Result represent result scored by the student.
type Result struct {
	Student     string    `json:"student"`
	SubmittedAt time.Time `json:"submitted_at"`
	Answers     []*Answer `json:"answers"`
	Percentage  int       `json:"percentage"`
	Points      int       `json:"points"`
	Total       int       `json:"total"`
}

// New checks solution submitted by student and returns Result.
func New(t *test.Test, s *test.Solution) *Result {
	r := Result{
		Student:     s.Student,
		SubmittedAt: s.SubmittedAt,
		Points:      0,
		Total:       t.TotalPoints(),
	}

	for i := range len(t.Tasks) {
		task := t.Tasks[i]
		answer := s.Answers[i]

		a := CheckAnswer(task, answer)

		if a.Correct {
			r.Points++
		}

		r.Answers = append(r.Answers, a)
	}

	r.Percentage = 100 * r.Points / r.Total

	return &r
}

// PerformanceCategory returns a number based on a scored percentage.
// 90 to 100 returns 0;
// 75 to 90 returns 1;
// 50 to 75 returns 2;
// 0 to 50 returns 3.
func (r Result) PerformanceCategory() int {
	if r.Percentage >= 90 {
		return 0
	}

	if r.Percentage >= 75 {
		return 1
	}

	if r.Percentage >= 50 {
		return 2
	}

	return 3
}
