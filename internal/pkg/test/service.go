package test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/shelepuginivan/hakutest/internal/config"
)

type TestService struct{}

func NewService() *TestService {
	return &TestService{}
}

func (s TestService) GetTestByName(name string) (Test, error) {
	test := Test{}
	testPath := s.GetTestPath(name)
	testFile, err := os.ReadFile(testPath)

	if err != nil {
		return test, err
	}

	err = json.Unmarshal(testFile, &test)

	return test, err
}

func (s TestService) GetTestByPath(path string) (Test, error) {
	test := Test{}
	testFile, err := os.ReadFile(path)

	if err != nil {
		return test, err
	}

	err = json.Unmarshal(testFile, &test)

	return test, err
}

func (s TestService) GetTestList() []string {
	testsDirectory := config.New().General.TestsDirectory
	testList := []string{}

	entries, err := os.ReadDir(testsDirectory)

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

func (s TestService) GetTestPath(name string) string {
	testsDirectory := config.New().General.TestsDirectory

	if !strings.HasSuffix(name, ".json") {
		name += ".json"
	}

	return filepath.Join(testsDirectory, name)
}

func (s TestService) Import(path string) error {
	name := filepath.Base(path)

	t, err := s.GetTestByPath(path)
	if err != nil {
		return err
	}

	return s.SaveToTestsDirectory(t, name)
}

func (s TestService) SaveToTestsDirectory(t Test, name string) error {
	testPath := s.GetTestPath(name)
	data, err := json.Marshal(t)

	if err != nil {
		return err
	}

	return os.WriteFile(testPath, data, 0666)
}
