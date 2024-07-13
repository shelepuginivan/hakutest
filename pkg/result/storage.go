package result

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/shelepuginivan/hakutest/internal/pkg/fsutil"
)

// GetForTest returns all results associated with the test.
//
// `testName` is a prefix directory in the results directory.
// For example, if `testName` "mytest" is provided, GetForTest would return
// results from `/results/dir/mytest` directory.
func GetForTest(testName string) (results []*Result, err error) {
	thisResultsDir := filepath.Join(resultsDirectory, testName)
	if !fsutil.DirExists(thisResultsDir) {
		return nil, fmt.Errorf("cannot find results for test: %s", testName)
	}

	entries, err := os.ReadDir(thisResultsDir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		entryPath := filepath.Join(thisResultsDir, entry.Name())
		entryContent, err := os.ReadFile(entryPath)
		if err != nil {
			continue
		}

		result := Result{}
		if err := json.Unmarshal(entryContent, &result); err != nil {
			continue
		}

		results = append(results, &result)
	}

	return results, nil
}

// AvailableResults returns a slice of available result prefix directories.
// Returned value can be used with method [GetForTest].
func AvailableResults() (names []string) {
	entries, err := os.ReadDir(resultsDirectory)
	if err != nil {
		return nil
	}

	for _, entry := range entries {
		if entry.IsDir() {
			names = append(names, entry.Name())
		}
	}

	return names
}

// Save saves result to the results directory.
//
// testName is a prefix directory in the results directory. For example, if it
// is equal to `mytest`, Save saves result to
// `/results/dir/mytest/<Result.Student>.json` file.
//
// Results are overwritten if and only if the configuration has field
// `overwrite_results` set to `true`.
func Save(r *Result, testName string) error {
	if strings.TrimSpace(testName) == "" {
		return fmt.Errorf("testName must be a valid directory name")
	}

	data, err := json.Marshal(r)
	if err != nil {
		return err
	}

	thisTestDir := filepath.Join(resultsDirectory, testName)

	if err = os.MkdirAll(thisTestDir, os.ModePerm|os.ModeDir); err != nil {
		return err
	}

	resultsFile := filepath.Join(thisTestDir, r.Student) + ".json"

	// Check whether result exists and `overwrite_results` is enabled.
	if !overwriteResults && fsutil.FileExists(resultsFile) {
		return nil
	}

	return os.WriteFile(resultsFile, data, os.ModePerm)
}
