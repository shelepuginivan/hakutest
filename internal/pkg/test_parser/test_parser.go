package parser

import (
	"encoding/json"
	"os"
	"path"
	"strings"

	"github.com/shelepuginivan/hakutest/internal/config"
)

func GetTestPath(name string) string {
	testsDirectory := config.Init().TestsDirectory

	if !strings.HasSuffix(name, ".json") {
		name += ".json"
	}

	return path.Join(testsDirectory, name)
}

func ParseTest(name string) (Test, error) {
	test := Test{}
	testPath := GetTestPath(name)
	testFile, err := os.ReadFile(testPath)

	if err != nil {
		return test, err
	}

	err = json.Unmarshal(testFile, &test)

	return test, err
}
