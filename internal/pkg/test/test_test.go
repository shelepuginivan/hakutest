package test

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/shelepuginivan/hakutest/internal/pkg/application"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/stretchr/testify/assert"
)

var app *application.App

func TestMain(m *testing.M) {
	defer setup()()
	m.Run()
}

func setup() func() {
	tmp, err := os.MkdirTemp("", "")
	if err != nil {
		panic(err)
	}

	app = &application.App{
		Config: &config.Config{
			General: &config.GeneralConfig{
				TestsDirectory: tmp,
			},
		},
	}

	testPath := filepath.Join(tmp, "__test__.json")

	mockTest := Test{
		Title: "Mock test",
	}

	data, err := json.Marshal(mockTest)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(testPath, data, 0666)
	if err != nil {
		panic(err)
	}

	return func() {
		if err := os.RemoveAll(tmp); err != nil {
			panic(err)
		}
	}
}

func TestTest_IsExpired(t *testing.T) {
	mockTest := Test{
		ExpiresIn: time.Now(),
	}

	assert.True(t, mockTest.IsExpired())
}

func TestTest_Sha256Sum(t *testing.T) {
	mockTest := Test{
		Title:  "Mock test",
		Author: "John Doe",
	}

	data, err := json.Marshal(mockTest)

	if err != nil {
		t.Fatal(err)
	}

	hasher := sha256.New()
	hasher.Write(data)

	assert.Equal(t, mockTest.Sha256Sum(), hex.EncodeToString(hasher.Sum(nil)))
}

func TestTestService_GetTestPath(t *testing.T) {
	s := NewService(application.New())
	testDir := config.New().General.TestsDirectory

	cases := []struct {
		name     string
		expected string
	}{
		{name: "a", expected: filepath.Join(testDir, "a.json")},
		{name: "a.json", expected: filepath.Join(testDir, "a.json")},
		{name: "some", expected: filepath.Join(testDir, "some.json")},
		{name: "longer name", expected: filepath.Join(testDir, "longer name.json")},
		{name: "", expected: filepath.Join(testDir, ".json")},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, s.GetTestPath(c.name))
	}
}

func TestTestService_GetTestByName(t *testing.T) {
	// get test created in setup function
	test, err := NewService(app).GetTestByName("__test__")

	assert.NoError(t, err)
	assert.Equal(t, "Mock test", test.Title)
}

func TestTestService_GetTestByPath(t *testing.T) {
	mockTest := Test{
		Title:       "Mock test",
		Author:      "John Doe",
		Target:      "Assert",
		Institution: "TestTestService_GetTestByName",
	}

	data, err := json.Marshal(mockTest)
	if err != nil {
		t.Fatal(err)
	}

	testFile, err := os.CreateTemp(os.TempDir(), "test")
	if err != nil {
		t.Fatal(err)
	}
	defer testFile.Close()

	testPath, err := filepath.Abs(testFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if _, err = testFile.Write(data); err != nil {
		t.Fatal(err)
	}

	test, err := NewService(app).GetTestByPath(testPath)

	assert.NoError(t, err)
	assert.Equal(t, mockTest.Title, test.Title)
	assert.Equal(t, mockTest.Author, test.Author)
	assert.Equal(t, mockTest.Target, test.Target)
	assert.Equal(t, mockTest.Institution, test.Institution)
}

func TestTestService_GetTestList(t *testing.T) {
	testList := NewService(app).GetTestList()

	assert.GreaterOrEqual(t, len(testList), 1)
	assert.Contains(t, testList, "__test__")
}

func TestTestService_SaveToTestsDirectory(t *testing.T) {
	mockTest := &Test{
		Title:       "Mock test",
		Author:      "John Doe",
		Target:      "Assert",
		Institution: "TestTestService_SaveToTestsDirectory",
	}

	err := NewService(app).SaveToTestsDirectory(mockTest, "__mock_test__")

	assert.NoError(t, err)

	test, err := NewService(app).GetTestByName("__mock_test__")

	assert.NoError(t, err)
	assert.EqualValues(t, mockTest, test)
}

func TestTestService_Import(t *testing.T) {
	mockTest := Test{
		Title:       "Mock test",
		Author:      "John Doe",
		Target:      "Assert",
		Institution: "TestTestService_Import",
	}

	data, err := json.Marshal(mockTest)
	if err != nil {
		t.Fatal(err)
	}

	testFile, err := os.CreateTemp(os.TempDir(), "test*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer testFile.Close()

	testPath, err := filepath.Abs(testFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if _, err = testFile.Write(data); err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, NewService(app).Import(testPath))

	test, err := NewService(app).GetTestByName(filepath.Base(testPath))

	assert.NoError(t, err)
	assert.Equal(t, mockTest.Title, test.Title)
	assert.Equal(t, mockTest.Author, test.Author)
	assert.Equal(t, mockTest.Target, test.Target)
	assert.Equal(t, mockTest.Institution, test.Institution)
}

func TestTestService_Remove(t *testing.T) {
	s := NewService(app)
	testName := "__TestService.Remove__"
	testPath := filepath.Join(app.Config.General.TestsDirectory, testName+".json")

	mockTest := Test{
		Title: "Mock test",
	}

	data, err := json.Marshal(mockTest)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(testPath, data, 0666)
	if err != nil {
		panic(err)
	}

	assert.NoError(t, s.Remove(testName))
	assert.NoFileExists(t, testPath)
}
