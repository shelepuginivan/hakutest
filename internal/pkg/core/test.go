package core

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"os"
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

func GetTest(name string) (Test, error) {
	test := Test{}
	testPath := GetTestPath(name)
	testFile, err := os.ReadFile(testPath)

	if err != nil {
		return test, err
	}

	err = json.Unmarshal(testFile, &test)

	return test, err
}

func (t Test) Save(name string) error {
	testPath := GetTestPath(name)
	data, err := json.Marshal(t)

	if err != nil {
		log.Fatal(err)
	}

	return os.WriteFile(testPath, data, 0666)
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
