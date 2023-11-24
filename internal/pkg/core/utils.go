package core

import (
	"encoding/json"
	"os"
	"path"
	"strings"

	"github.com/shelepuginivan/hakutest/internal/config"
)

func CompareAnswers(received, expected string) bool {
	return strings.TrimSpace(strings.ToLower(received)) == strings.TrimSpace(strings.ToLower(expected))
}

func GetTestPath(name string) string {
	testsDirectory := config.Init().General.TestsDirectory

	if !strings.HasSuffix(name, ".json") {
		name += ".json"
	}

	return path.Join(testsDirectory, name)
}

func Import(file string) error {
	testFile, err := os.ReadFile(file)
	testPath := GetTestPath(path.Base(file))
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
