package parser

import (
	"encoding/json"
	"os"
	"path"
	"strings"
	"time"

	"github.com/shelepuginivan/hakutest/internal/config"
)

type Test struct {
	Title       string    `json:"title"`
	Target      string    `json:"target"`
	Subject     string    `json:"subject"`
	Author      string    `json:"author"`
	Institution string    `json:"institution"`
	CreatedAt   time.Time `json:"createdAt"`
	ExpiresIn   time.Time `json:"expiresIn"`
	Tasks       []struct {
		Type       string   `json:"type"`
		Text       string   `json:"text"`
		Attachment string   `json:"attachment"`
		Options    []string `json:"options"`
	} `json:"tasks"`
}

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
