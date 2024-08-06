package result

import (
	"path/filepath"
	"strings"

	"github.com/shelepuginivan/fsutil"
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
		Type:  task.Type,
		Value: answer,
	}

	if task.Type == "detailed" {
		a.Correct = len(NormalizeAnswer(answer)) > 0
	} else {
		a.Correct = NormalizeAnswer(answer) == NormalizeAnswer(task.Answer)
	}

	return &a
}

// DeleteMany deletes results by names. The returned value is the number of
// successfully deleted result subdirectories.
func DeleteMany(names ...string) (deleted int) {
	for _, name := range names {
		err := fsutil.RemoveAllIfExists(filepath.Join(resultsDirectory, name))

		if err != nil {
			continue
		}

		deleted++
	}

	return deleted
}
