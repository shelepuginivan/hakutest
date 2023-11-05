package parser

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path"
	"strings"

	"github.com/shelepuginivan/hakutest/internal/config"
)

func GetTestPath(name string) string {
	testsDirectory := config.Init().General.TestsDirectory

	if !strings.HasSuffix(name, ".json") {
		name += ".json"
	}

	return path.Join(testsDirectory, name)
}

func GetTestCheckSum(name string) (string, error) {
	hasher := sha256.New()
	testPath := GetTestPath(name)
	test, err := os.ReadFile(testPath)

	if err != nil {
		return "", err
	}

	hasher.Write(test)

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
