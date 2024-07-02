// Package test provides functionality to manipulate tests.
package test

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
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
	Title       string    `json:"title"`       // Title of the test.
	Target      string    `json:"target"`      // Target audience of the test.
	Description string    `json:"description"` // Description of the test.
	Subject     string    `json:"subject"`     // Subject of the test.
	Author      string    `json:"author"`      // Author of the test.
	Institution string    `json:"institution"` // Institution associated with the test.
	CreatedAt   time.Time `json:"createdAt"`   // Creation time of the test.
	ExpiresAt   time.Time `json:"expiresAt"`   // Expiration time of the test.
	Tasks       []*Task   `json:"tasks"`       // Tasks of the test.
}

// IsExpired reports whether the test is expired.
// If the ExpiresIn field is zero, it returns true.
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
