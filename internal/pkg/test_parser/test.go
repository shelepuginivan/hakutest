package parser

import (
	"encoding/json"
	"os"
	"path"
	"time"
)

type Attachment struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Src  string `json:"src"`
}

type Task struct {
	Type       string     `json:"type"`
	Text       string     `json:"text"`
	Attachment Attachment `json:"attachment"`
	Options    []string   `json:"options"`
	Answer     string     `json:"answer"`
}

type Test struct {
	Title       string    `json:"title"`
	Target      string    `json:"target"`
	Description string    `json:"description"`
	Subject     string    `json:"subject"`
	Author      string    `json:"author"`
	Institution string    `json:"institution"`
	CreatedAt   time.Time `json:"createdAt"`
	ExpiresIn   time.Time `json:"expiresIn"`
	Tasks       []Task    `json:"tasks"`
}

func Get(name string) (Test, error) {
	test := Test{}
	testPath := GetTestPath(name)
	testFile, err := os.ReadFile(testPath)

	if err != nil {
		return test, err
	}

	err = json.Unmarshal(testFile, &test)

	return test, err
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
