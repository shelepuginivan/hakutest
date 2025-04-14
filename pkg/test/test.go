// Package test provides functionality to manipulate tests.
package test

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"math/rand/v2"
	"time"
)

const (
	TaskSingle   = "single"   // Task with a single answer.
	TaskMultiple = "multiple" // Task with multiple answers.
	TaskOpen     = "open"     // Open-ended task.
	TaskDetailed = "detailed" // Task with detailed answer.
)

// Task represents a Test task with a type, text, attachment, options, and answer.
type Task struct {
	Type    string   `json:"type"`    // Type of the task.
	Text    string   `json:"text"`    // Text of the task, usually its terms.
	Options []string `json:"options"` // Answer options of the task.
	Answer  string   `json:"answer"`  // Correct answer of the task, zero-indexed.
}

// Test represents a test with a title, target, description, subject, author, institution, creation date, expiration date, and tasks.
type Test struct {
	Title        string    `json:"title"`        // Title of the test.
	Target       string    `json:"target"`       // Target audience of the test.
	Description  string    `json:"description"`  // Description of the test.
	Subject      string    `json:"subject"`      // Subject of the test.
	Author       string    `json:"author"`       // Author of the test.
	Institution  string    `json:"institution"`  // Institution associated with the test.
	CreatedAt    time.Time `json:"createdAt"`    // Creation time of the test.
	ExpiresAt    time.Time `json:"expiresAt"`    // Expiration time of the test.
	Tasks        []*Task   `json:"tasks"`        // Tasks of the test.
	ShuffleTasks bool      `json:"shuffleTasks"` // Whether display tasks in random order.
}

// TotalPoints returns total points of the test.
func (t Test) TotalPoints() int {
	return len(t.Tasks)
}

// IsExpired reports whether the test is expired.
// If the ExpiresIn field is zero, it returns false.
func (t Test) IsExpired() bool {
	return !t.ExpiresAt.IsZero() && t.ExpiresAt.Before(time.Now())
}

// Sha256Sum returns the sha256 checksum of the test.
func (t Test) Sha256Sum() string {
	hasher := sha256.New()
	data, err := json.Marshal(t)

	if err != nil {
		return ""
	}

	hasher.Write(data)

	return hex.EncodeToString(hasher.Sum(nil))
}

// TaskKeys returns slice of indices of tasks, representing order in which
// clients should display tasks. The order depends on [Test.ShuffleTasks].
func (t Test) TaskKeys() []int {
	length := len(t.Tasks)

	keys := make([]int, length)
	for i := range length {
		keys[i] = i
	}

	if t.ShuffleTasks {
		rand.Shuffle(length, func(i, j int) {
			keys[i], keys[j] = keys[j], keys[i]
		})
	}

	return keys
}
