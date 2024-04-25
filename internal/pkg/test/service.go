package test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/shelepuginivan/hakutest/internal/pkg/application"
)

// TestService is a struct that provides methods for manipulating Test structures.
type TestService struct {
	app *application.App
}

// NewService returns a TestService instance.
func NewService(app *application.App) *TestService {
	return &TestService{app: app}
}

// GetTestByName retrieves the Test from the directory specified in the configuration by the test name.
func (s TestService) GetTestByName(name string) (*Test, error) {
	test := &Test{}
	testPath := s.GetTestPath(name)
	testFile, err := os.ReadFile(testPath)

	if err != nil {
		return test, err
	}

	err = json.Unmarshal(testFile, &test)

	return test, err
}

// GetTestByPath retrieves the Test by its file path.
func (s TestService) GetTestByPath(path string) (*Test, error) {
	test := &Test{}
	testFile, err := os.ReadFile(path)

	if err != nil {
		return test, err
	}

	err = json.Unmarshal(testFile, &test)

	return test, err
}

// GetTestList retrieves a list of test names from the tests directory specified in the configuration.
func (s TestService) GetTestList() []string {
	testList := []string{}

	entries, err := os.ReadDir(s.app.Config.General.TestsDirectory)

	if err != nil {
		return testList
	}

	for _, file := range entries {
		testName := file.Name()

		if !file.IsDir() && strings.HasSuffix(testName, ".json") {
			testList = append(testList, strings.TrimSuffix(testName, ".json"))
		}
	}

	return testList
}

// GetTestPath returns the absolute path of the Test by its name.
// It assumes that the test is stored in the tests directory specified in the configuration.
// It doesn't check whether a test with this name exists.
func (s TestService) GetTestPath(name string) string {
	if !strings.HasSuffix(name, ".json") {
		name += ".json"
	}

	return filepath.Join(s.app.Config.General.TestsDirectory, name)
}

// Import copies the test file saved by path to the tests directory defined in the configuration.
func (s TestService) Import(path string) error {
	name := filepath.Base(path)

	t, err := s.GetTestByPath(path)
	if err != nil {
		return err
	}

	return s.SaveToTestsDirectory(t, name)
}

// SaveToTestsDirectory saves the Test as a JSON file in the tests directory specified in the configuration.
// The name is used as a filename.
func (s TestService) SaveToTestsDirectory(t *Test, name string) error {
	testPath := s.GetTestPath(name)
	data, err := json.Marshal(t)

	if err != nil {
		return err
	}

	return os.WriteFile(testPath, data, 0666)
}

// Remove removes the file associated with the test by its name.
// It assumes that the test is stored in the tests directory specified in the configuration.
func (s TestService) Remove(testName string) error {
	return os.Remove(s.GetTestPath(testName))
}
