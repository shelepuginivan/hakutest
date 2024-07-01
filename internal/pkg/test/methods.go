package test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/shelepuginivan/hakutest/internal/pkg/fsutil"
)

// NormalizeName appends suffix `.json` to the test name if it is missing.
func NormalizeName(name string) string {
	if strings.HasSuffix(name, ".json") {
		return name
	}

	return fmt.Sprintf("%s.json", name)
}

// GetList returns a slice of names of tests stored in the tests directory.
func GetList() []string {
	var tests []string

	files, err := os.ReadDir(testsDirectory)
	if err != nil {
		return tests
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			tests = append(tests, file.Name())
		}
	}

	return tests
}

// GetByName returns test by name.
// If test does not exist, error is returned.
func GetByName(name string) (*Test, error) {
	var test Test
	testPath := filepath.Join(testsDirectory, NormalizeName(name))

	if !fsutil.FileExists(testPath) {
		return nil, fmt.Errorf("test %s not found", name)
	}

	data, err := os.ReadFile(testPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &test)
	if err != nil {
		return nil, err
	}

	return &test, nil
}
