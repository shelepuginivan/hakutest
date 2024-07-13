package test

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
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

// PrettifyName removes suffix `.json` from the test name.
func PrettifyName(name string) string {
	return strings.TrimSuffix(name, ".json")
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
			tests = append(tests, PrettifyName(file.Name()))
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

// Import imports test to the tests directory. If test with the same title
// exists, append a numeric suffix. The suffix is incremented until test name
// is unique.
func Import(test []byte) error {
	var t Test
	if err := json.Unmarshal(test, &t); err != nil {
		return err
	}

	if strings.TrimSpace(t.Title) == "" {
		return fmt.Errorf("title must not be empty")
	}

	testPath := filepath.Join(testsDirectory, NormalizeName(t.Title))

	if !fsutil.FileExists(testPath) {
		return fsutil.WriteAll(testPath, test)
	}

	for i := 1; true; i++ {
		testName := fmt.Sprintf("%s (%d)", t.Title, i)
		testPath = filepath.Join(testsDirectory, NormalizeName(testName))

		if !fsutil.FileExists(testPath) {
			return fsutil.WriteAll(testPath, test)
		}
	}

	return nil
}

// Save saves test to the tests directory. Unlike [Import], if test with the
// same title exists, it is overwritten.
func Save(t *Test) error {
	if strings.TrimSpace(t.Title) == "" {
		return fmt.Errorf("title must not be empty")
	}
	testPath := filepath.Join(testsDirectory, NormalizeName(t.Title))

	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	if !fsutil.FileExists(testPath) {
		return fsutil.WriteAll(testPath, data)
	}

	return fsutil.WriteAll(testPath, data)
}

// WriteJSON writes raw test JSON file to w.
func WriteJSON(w io.Writer, name string) error {
	t, err := GetByName(name)
	if err != nil {
		return err
	}

	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

// WriteZip appends tests by names to a zip archive and writes the resulting
// archive to w.
func WriteZip(w io.Writer, names ...string) error {
	zipWriter := zip.NewWriter(w)

	for _, name := range names {
		t, err := GetByName(name)
		if err != nil {
			continue
		}

		data, err := json.Marshal(t)
		if err != nil {
			continue
		}

		f, err := zipWriter.Create(NormalizeName(name))
		if err != nil {
			continue
		}

		f.Write(data)
	}

	return zipWriter.Close()
}

// DeleteMany deletes tests by names. The returned value is the number of
// successfully deleted test files.
func DeleteMany(names ...string) (deleted int) {
	for _, name := range names {
		err := os.Remove(filepath.Join(testsDirectory, NormalizeName(name)))

		if err != nil {
			continue
		}

		deleted++
	}

	return deleted
}
