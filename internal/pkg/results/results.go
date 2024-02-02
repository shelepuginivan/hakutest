// Package results provides functionality to manipulate test results.
package results

import "time"

// TestResults represents a result of the test.
// It contains name of the student, submission time, the results themselves, and information about the test.
type TestResults struct {
	Student     string    `yaml:"student"`     // Name of the student.
	SubmittedAt time.Time `yaml:"submittedAt"` // Submission time.
	Results     Results   `yaml:"results"`     // Results.
	Test        TestInfo  `yaml:"test"`        // Information about the test.
}

// TaskResult represents a solution of each task.
type TaskResult struct {
	Answer  string `yaml:"answer"`  // Student answer.
	Correct bool   `yaml:"correct"` // The correctness of the answer.
}

// Results represents information about the student performance.
// It contains scored points, total points of the test, scored percentage, and solutions of every task of the test.
type Results struct {
	Points     int                   `yaml:"points"`     // Points scored by student.
	Total      int                   `yaml:"total"`      // Total points of the test.
	Percentage int                   `yaml:"percentage"` // Percentage scored by student.
	Tasks      map[string]TaskResult `yaml:"tasks"`      // Solutions of the tasks.
}

// TestInfo represents information about the test.
// It contains title of the test, its author, and sha256 checksum.
type TestInfo struct {
	Title  string `yaml:"title"`  // Title of the test.
	Author string `yaml:"author"` // Author of the test.
	Sha256 string `yaml:"sha256"` // Checksum of the test.
}
