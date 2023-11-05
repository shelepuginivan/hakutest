package parser

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
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

func (t Test) Sha256Sum() string {
	hasher := sha256.New()
	data, err := json.Marshal(t)

	if err != nil {
		return ""
	}

	hasher.Write(data)

	return hex.EncodeToString(hasher.Sum(nil))
}
