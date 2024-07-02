package test

import "time"

// Solution represents test solution submitted by the student.
type Solution struct {
	Student     string    // Student who submitted the solution.
	Answers     []string  // Answers to the tasks.
	SubmittedAt time.Time // Submission time.
}
