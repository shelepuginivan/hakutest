package core

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/shelepuginivan/hakutest/internal/config"
)

func CompareAnswers(received, expected string) bool {
	return strings.TrimSpace(strings.ToLower(received)) == strings.TrimSpace(strings.ToLower(expected))
}

func GetTestList() []string {
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

func GetTestPath(name string) string {
	testsDirectory := config.New().General.TestsDirectory

	if !strings.HasSuffix(name, ".json") {
		name += ".json"
	}

	return filepath.Join(testsDirectory, name)
}

func Import(file string) error {
	testFile, err := os.ReadFile(file)
	testPath := GetTestPath(filepath.Base(file))
	test := Test{}

	if err != nil {
		return err
	}

	err = json.Unmarshal(testFile, &test)

	if err != nil {
		return err
	}

	return os.WriteFile(testPath, testFile, 0666)
}
