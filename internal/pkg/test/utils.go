package test

import (
	"path/filepath"
	"strings"

	"github.com/shelepuginivan/hakutest/internal/config"
)

func GetTestPath(name string) string {
	testsDirectory := config.New().General.TestsDirectory

	if !strings.HasSuffix(name, ".json") {
		name += ".json"
	}

	return filepath.Join(testsDirectory, name)
}
