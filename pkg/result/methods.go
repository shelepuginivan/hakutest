package result

import (
	"strings"

	"github.com/shelepuginivan/hakutest/pkg/test"
)

// NormalizeAnswer returns answer in normal form.
// It trims leading and trailing whitespaces, and converts it to lowercase.
func NormalizeAnswer(answer string) string {
	return strings.ToLower(strings.TrimSpace(answer))
}

// CheckAnswer checks whether answer is correct and returns Answer struct.
// Correctness of answer depends on type of the task.
func CheckAnswer(task *test.Task, answer string) *Answer {
	a := Answer{
		Value: answer,
	}

	if task.Type == "detailed" {
		a.Correct = len(NormalizeAnswer(answer)) > 0
	} else {
		a.Correct = NormalizeAnswer(answer) == NormalizeAnswer(task.Answer)
	}

	return &a
}
