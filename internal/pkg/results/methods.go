package results

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/shelepuginivan/hakutest/internal/pkg/fsutil"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
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

// Save saves result to the results directory.
func Save(r *Result, testName string) error {
	data, err := json.Marshal(r)
	if err != nil {
		return err
	}

	thisTestDir := filepath.Join(resultsDirectory, testName)

	if err = os.MkdirAll(thisTestDir, os.ModePerm|os.ModeDir); err != nil {
		return err
	}

	resultsFile := filepath.Join(thisTestDir, r.Student) + ".json"

	// Check whether result exist and overwrite is enabled.
	if !overwriteResults && fsutil.FileExists(resultsFile) {
		return nil
	}

	return os.WriteFile(resultsFile, data, os.ModePerm)
}
