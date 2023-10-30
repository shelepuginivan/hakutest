package parser

import (
	"encoding/json"
	"os"
	"path"
	"strings"

	"github.com/shelepuginivan/hakutest/internal/config"
)

func ParseTest(name string) (Test, error) {
	if !strings.HasSuffix(name, ".json") {
		name += ".json"
	}

	test := Test{}
	testsDirectory := config.Init().TestsDirectory
	testPath := path.Join(testsDirectory, name)
	testFile, err := os.ReadFile(testPath)

	if err != nil {
		return test, err
	}

	err = json.Unmarshal(testFile, &test)

	return test, err
}
