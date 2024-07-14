package test_test

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/shelepuginivan/hakutest/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockWriter struct {
	mock.Mock
}

func (m *MockWriter) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("write failed: mock writer")
}

type TestCase struct {
	Test     *test.Test
	Expected any
	Actual   any
}

func TestTest_TotalPoints(t *testing.T) {
	cases := []TestCase{
		{
			Test:     &test.Test{Tasks: []*test.Task{}},
			Expected: 0,
		},
		{
			Test:     &test.Test{Tasks: []*test.Task{{}, {}, {}}},
			Expected: 3,
		},
	}

	for _, c := range cases {
		c.Actual = c.Test.TotalPoints()
		assert.Equal(t, c.Expected, c.Actual)
	}
}

func TestTest_IsExpired(t *testing.T) {
	cases := []TestCase{
		{
			Test:     &test.Test{ExpiresAt: time.Now().Add(time.Hour)},
			Expected: false,
		},
		{
			Test:     &test.Test{ExpiresAt: time.Now().Add(-time.Hour)},
			Expected: true,
		},
		{
			Test:     &test.Test{},
			Expected: false,
		},
	}

	for _, c := range cases {
		c.Actual = c.Test.IsExpired()
		assert.Equal(t, c.Expected, c.Actual)
	}
}

func TestTest_Sha256Sum(t *testing.T) {
	tst := test.Test{
		Title:       "title",
		Description: "description",
		Author:      "go",
		Subject:     "hakutest",
		Target:      "go test ./...",
		Institution: "-",
		ExpiresAt:   time.Now().Add(time.Hour * 72),
		Tasks: []*test.Task{
			{
				Type:    "single",
				Text:    "# my task",
				Answer:  "1",
				Options: []string{"0", "1", "2", "3"},
			},
		},
	}

	hasher := sha256.New()
	data, _ := json.Marshal(tst)
	hasher.Write(data)

	assert.Equal(t, hex.EncodeToString(hasher.Sum(nil)), tst.Sha256Sum())
}

func TestNormalizeName(t *testing.T) {
	cases := []TestCase{
		{
			Expected: "test.json",
			Actual:   test.NormalizeName("test"),
		},
		{
			Expected: "test name with spaces.json",
			Actual:   test.NormalizeName("test name with spaces"),
		},
		{
			Expected: "test with suffix (2).json",
			Actual:   test.NormalizeName("test with suffix (2)"),
		},
		{
			Expected: "already has ext.json",
			Actual:   test.NormalizeName("already has ext.json"),
		},
		{
			Expected: ".json",
			Actual:   test.NormalizeName(""),
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.Expected, c.Actual)
	}
}

func TestPrettifyName(t *testing.T) {
	cases := []TestCase{
		{
			Expected: "test",
			Actual:   test.PrettifyName("test.json"),
		},
		{
			Expected: "test with spaces",
			Actual:   test.PrettifyName("test with spaces.json"),
		},
		{
			Expected: "with suffix",
			Actual:   test.PrettifyName("with suffix.json"),
		},
		{
			Expected: "already pretty",
			Actual:   test.PrettifyName("already pretty"),
		},
		{
			Expected: "",
			Actual:   test.PrettifyName(".json"),
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.Expected, c.Actual)
	}
}

func TestGetList(t *testing.T) {
	t.Run("should return test names", func(t *testing.T) {
		tmp := t.TempDir()
		test.Init(test.Config{
			Path: tmp,
		})

		os.WriteFile(filepath.Join(tmp, "some.json"), []byte{}, os.ModePerm)
		os.WriteFile(filepath.Join(tmp, "another.json"), []byte{}, os.ModePerm)
		os.WriteFile(filepath.Join(tmp, "should ignore non-json.txt"), []byte{}, os.ModePerm)
		os.Mkdir(filepath.Join(tmp, "should ignore directories"), os.ModePerm|os.ModeDir)

		l := test.GetList()

		assert.Len(t, l, 2)
		assert.Contains(t, l, "some")
		assert.Contains(t, l, "another")
		assert.NotContains(t, l, "should ignore non-json")
		assert.NotContains(t, l, "should ignore non-json.txt")
		assert.NotContains(t, l, "should ignore directories")
	})

	t.Run("should return empty slice if error occurres", func(t *testing.T) {
		test.Init(test.Config{
			Path: "this directory does not exist",
		})

		assert.Empty(t, test.GetList())
	})
}

func TestGetByName(t *testing.T) {
	expected := &test.Test{
		Title: "my new test",
	}
	data, err := json.Marshal(expected)
	if err != nil {
		panic(err)
	}

	tmp := t.TempDir()
	test.Init(test.Config{
		Path: tmp,
	})

	t.Run("should get test by name", func(t *testing.T) {
		os.WriteFile(filepath.Join(tmp, "some.json"), data, os.ModePerm)

		actual, err := test.GetByName("some")
		assert.NoError(t, err)
		assert.EqualValues(t, expected, actual)
	})

	t.Run("should return error if test does not exist", func(t *testing.T) {
		_, err := test.GetByName("this one does not exist")
		assert.Error(t, err)
	})

	t.Run("should return error if test cannot be read", func(t *testing.T) {
		os.WriteFile(filepath.Join(tmp, "write only.json"), data, 0333)

		_, err := test.GetByName("write only")
		assert.Error(t, err)
	})

	t.Run("should return error if test is not a valid JSON", func(t *testing.T) {
		os.WriteFile(filepath.Join(tmp, "not a json.json"), []byte("at least not a valid one"), os.ModePerm)

		_, err := test.GetByName("not a json")
		assert.Error(t, err)
	})
}

func TestImport(t *testing.T) {
	tmp := t.TempDir()
	test.Init(test.Config{
		Path: tmp,
	})

	t.Run("should import test", func(t *testing.T) {
		tst := &test.Test{
			Title: "my new test",
		}
		data, _ := json.Marshal(tst)

		assert.NoError(t, test.Import(data))
		assert.Contains(t, test.GetList(), "my new test")
	})

	t.Run("should append numeric suffix to avoid duplications", func(t *testing.T) {
		tst := &test.Test{
			Title: "my new test",
		}
		data, _ := json.Marshal(tst)

		assert.NoError(t, test.Import(data))
		assert.Contains(t, test.GetList(), "my new test")
		assert.Contains(t, test.GetList(), "my new test (1)")
	})

	t.Run("should return error if test has no title", func(t *testing.T) {
		tst := &test.Test{
			Description: "it does not contain a title",
		}
		data, err := json.Marshal(tst)
		if err != nil {
			panic(err)
		}

		assert.Error(t, test.Import(data))
	})

	t.Run("should return error if test is not a valid JSON", func(t *testing.T) {
		data := []byte("this is not a JSON")

		assert.Error(t, test.Import(data))
	})
}

func TestSave(t *testing.T) {
	tmp := t.TempDir()
	test.Init(test.Config{
		Path: tmp,
	})

	t.Run("should import test", func(t *testing.T) {
		tst := &test.Test{
			Title: "my new test",
		}

		assert.NoError(t, test.Save(tst))
		assert.Contains(t, test.GetList(), "my new test")
	})

	t.Run("should not append numeric suffix and rewrite existing test", func(t *testing.T) {
		tst := &test.Test{
			Title:       "my new test",
			Description: "a new description",
		}

		assert.NoError(t, test.Save(tst))
		assert.Contains(t, test.GetList(), "my new test")
		assert.NotContains(t, test.GetList(), "my new test (1)")

		actual, err := test.GetByName("my new test")

		assert.NoError(t, err)
		assert.Equal(t, tst, actual)
	})

	t.Run("should return error if test has no title", func(t *testing.T) {
		tst := &test.Test{Description: "it does not contain a title"}
		assert.Error(t, test.Save(tst))
	})
}
func TestWriteJSON(t *testing.T) {
	expected := &test.Test{
		Title: "my new test",
		Tasks: []*test.Task{
			{
				Type:    "single",
				Text:    "# Lorem ipsum\n\r\n\rDolor sit amet",
				Options: []string{"ok", "no"},
				Answer:  "0",
			},
		},
	}
	data, _ := json.Marshal(expected)

	tmp := t.TempDir()
	test.Init(test.Config{
		Path: tmp,
	})

	os.WriteFile(filepath.Join(tmp, "some.json"), data, os.ModePerm)

	t.Run("should write tests in JSON format", func(t *testing.T) {
		w := bytes.NewBuffer(nil)
		err := test.WriteJSON(w, "some")

		assert.NoError(t, err)
		assert.Equal(t, data, w.Bytes())
	})

	t.Run("should return error if test does not exist", func(t *testing.T) {
		err := test.WriteJSON(bytes.NewBuffer(nil), "this test does not exist")
		assert.Error(t, err)
	})

	t.Run("should return error if write fails", func(t *testing.T) {
		err := test.WriteJSON(&MockWriter{}, "this test does not exist")
		assert.Error(t, err)
	})
}

func TestWriteZip(t *testing.T) {
	tmp := t.TempDir()
	test.Init(test.Config{
		Path: tmp,
	})

	t1 := []byte("{\"title\":\"test 1\"}")
	t2 := []byte("{\"title\":\"test 2\"}")
	t3 := []byte("{\"title\":\"test 3\"}")
	os.WriteFile(filepath.Join(tmp, "test 1.json"), t1, os.ModePerm)
	os.WriteFile(filepath.Join(tmp, "test 2.json"), t2, os.ModePerm)
	os.WriteFile(filepath.Join(tmp, "test 3.json"), t3, os.ModePerm)

	b := bytes.NewBuffer(nil)
	test.WriteZip(b, "test 1", "test 2", "test 3", "does not exist")

	zipReader, err := zip.NewReader(bytes.NewReader(b.Bytes()), int64(b.Len()))
	assert.NoError(t, err)

	var fileNames []string
	for _, file := range zipReader.File {
		fileNames = append(fileNames, file.Name)
	}

	assert.Len(t, fileNames, 3)
	assert.Contains(t, fileNames, "test 1.json")
	assert.Contains(t, fileNames, "test 2.json")
	assert.Contains(t, fileNames, "test 3.json")
	assert.NotContains(t, fileNames, "does not exist.json")
}

func TestDeleteMany(t *testing.T) {
	tmp := t.TempDir()
	test.Init(test.Config{
		Path: tmp,
	})

	os.WriteFile(filepath.Join(tmp, "1.json"), []byte{}, os.ModePerm)
	os.WriteFile(filepath.Join(tmp, "2.json"), []byte{}, os.ModePerm)
	os.WriteFile(filepath.Join(tmp, "3.json"), []byte{}, os.ModePerm)
	os.WriteFile(filepath.Join(tmp, "4.json"), []byte{}, os.ModePerm)
	os.WriteFile(filepath.Join(tmp, "5.json"), []byte{}, os.ModePerm)

	assert.Equal(t, 1, test.DeleteMany("4"))
	assert.Equal(t, 3, test.DeleteMany("1", "2", "3"))
	assert.Equal(t, 0, test.DeleteMany())
	assert.Equal(t, 0, test.DeleteMany("these", "don't", "exist"))
	assert.Contains(t, test.GetList(), "5")
}
