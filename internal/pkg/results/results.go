// Package results provides results management for the app.
package results

import "github.com/shelepuginivan/hakutest/internal/pkg/test"

// Answer represents answer given by the student.
type Answer struct {
	Value   string `json:"value"`   // The answer.
	Correct bool   `json:"correct"` // Whether answer is correct.
}

// Result represent result scored by the student.
type Result struct {
	Student    string    `json:"student"`
	Answers    []*Answer `json:"answers"`
	Percentage int       `json:"percentage"`
	Points     int       `json:"points"`
	Total      int       `json:"total"`
}

// New checks solution submitted by student and returns Result.
func New(t *test.Test, s *test.Solution) *Result {
	r := Result{
		Student: s.Student,
		Points:  0,
		Total:   t.TotalPoints(),
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
