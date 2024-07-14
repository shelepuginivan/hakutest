package result_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/shelepuginivan/hakutest/pkg/result"
	"github.com/shelepuginivan/hakutest/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestResult_PerformanceCategory(t *testing.T) {
	cases := []struct {
		res      *result.Result
		expected int
	}{
		{
			res:      &result.Result{Percentage: 100},
			expected: 0,
		},
		{
			res:      &result.Result{Percentage: 90},
			expected: 0,
		},
		{
			res:      &result.Result{Percentage: 89},
			expected: 1,
		},
		{
			res:      &result.Result{Percentage: 75},
			expected: 1,
		},
		{
			res:      &result.Result{Percentage: 74},
			expected: 2,
		},
		{
			res:      &result.Result{Percentage: 50},
			expected: 2,
		},
		{
			res:      &result.Result{Percentage: 49},
			expected: 3,
		},
		{
			res:      &result.Result{Percentage: 0},
			expected: 3,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, c.res.PerformanceCategory())
	}
}

func TestNew(t *testing.T) {
	t.Run("should pass test 1", func(t *testing.T) {
		submittedAt := time.Now()

		tst := &test.Test{
			Tasks: []*test.Task{
				{
					Type:   test.TaskSingle,
					Answer: "1",
				},
				{
					Type:   test.TaskMultiple,
					Answer: "0,1,2,3",
				},
				{
					Type:   test.TaskOpen,
					Answer: "some",
				},
				{
					Type: test.TaskDetailed,
				},
			},
		}
		s := &test.Solution{
			Student:     "someone",
			SubmittedAt: submittedAt,
			Answers:     []string{"1", "0,1,2,3", "some", "detailed"},
		}

		expected := &result.Result{
			Student:     "someone",
			SubmittedAt: submittedAt,
			Points:      4,
			Total:       4,
			Percentage:  100,
			Answers: []*result.Answer{
				{
					Type:    test.TaskSingle,
					Value:   "1",
					Correct: true,
				},
				{
					Type:    test.TaskMultiple,
					Value:   "0,1,2,3",
					Correct: true,
				},
				{
					Type:    test.TaskOpen,
					Value:   "some",
					Correct: true,
				},
				{
					Type:    test.TaskDetailed,
					Value:   "detailed",
					Correct: true,
				},
			},
		}
		actual := result.New(tst, s)

		assert.True(t, reflect.DeepEqual(expected, actual))
	})

	t.Run("should pass test 2", func(t *testing.T) {
		submittedAt := time.Now()

		tst := &test.Test{
			Tasks: []*test.Task{
				{
					Type:   test.TaskSingle,
					Answer: "1",
				},
			},
		}
		s := &test.Solution{
			Student:     "someone",
			SubmittedAt: submittedAt,
			Answers:     []string{"2"},
		}

		expected := &result.Result{
			Student:     "someone",
			SubmittedAt: submittedAt,
			Points:      0,
			Total:       1,
			Percentage:  0,
			Answers: []*result.Answer{
				{
					Type:    test.TaskSingle,
					Value:   "2",
					Correct: false,
				},
			},
		}
		actual := result.New(tst, s)

		assert.True(t, reflect.DeepEqual(expected, actual))
	})
}

func TestNormalizeAnswer(t *testing.T) {
	cases := []struct {
		a        string
		expected string
	}{
		{a: "normalized", expected: "normalized"},
		{a: " normalized ", expected: "normalized"},
		{a: "  1,2,3  ", expected: "1,2,3"},
		{a: "   a long, detailed answer\nyes\n\t", expected: "a long, detailed answer\nyes"},
		{a: "", expected: ""},
		{a: "AAAAAAAAAAAA", expected: "aaaaaaaaaaaa"},
		{a: " now WITH spaces AND uppercase ", expected: "now with spaces and uppercase"},
		{a: "\t\t\t\n\n\t\r\t\n\n\r\t", expected: ""},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, result.NormalizeAnswer(c.a))
	}
}

func TestCheckAnswer(t *testing.T) {
	cases := []struct {
		ans      string
		task     *test.Task
		expected *result.Answer
	}{
		{
			ans: "3",
			task: &test.Task{
				Type:   "single",
				Answer: "3",
			},
			expected: &result.Answer{
				Type:    "single",
				Value:   "3",
				Correct: true,
			},
		},
		{
			ans: "5",
			task: &test.Task{
				Type:   "single",
				Answer: "2",
			},
			expected: &result.Answer{
				Type:    "single",
				Value:   "5",
				Correct: false,
			},
		},
		{
			ans: "1,2,3",
			task: &test.Task{
				Type:   "multiple",
				Answer: "1,2,3",
			},
			expected: &result.Answer{
				Type:    "multiple",
				Value:   "1,2,3",
				Correct: true,
			},
		},
		{
			ans: "1,2,3,4,5",
			task: &test.Task{
				Type:   "multiple",
				Answer: "3,4",
			},
			expected: &result.Answer{
				Type:    "multiple",
				Value:   "1,2,3,4,5",
				Correct: false,
			},
		},
		{
			ans: "some text",
			task: &test.Task{
				Type:   "open",
				Answer: "some text",
			},
			expected: &result.Answer{
				Type:    "open",
				Value:   "some text",
				Correct: true,
			},
		},
		{
			ans: "wrong answer",
			task: &test.Task{
				Type:   "open",
				Answer: "indeed",
			},
			expected: &result.Answer{
				Type:    "open",
				Value:   "wrong answer",
				Correct: false,
			},
		},
		{
			ans: "This is a detailed answer.\nIt is to be checked manually, but it is marked as correct",
			task: &test.Task{
				Type: "detailed",
			},
			expected: &result.Answer{
				Type:    "detailed",
				Value:   "This is a detailed answer.\nIt is to be checked manually, but it is marked as correct",
				Correct: true,
			},
		},
		{
			ans: "",
			task: &test.Task{
				Type: "detailed",
			},
			expected: &result.Answer{
				Type:    "detailed",
				Value:   "",
				Correct: false,
			},
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, result.CheckAnswer(c.task, c.ans))
	}
}

func TestGetForTest(t *testing.T) {
	tmp := t.TempDir()
	result.Init(result.Config{
		Path:      tmp,
		Overwrite: false,
	})

	t.Run(
		"should get slice of valid results, ignore directories and invalid entries",
		func(t *testing.T) {
			prefixDir := filepath.Join(tmp, "prefix")

			os.MkdirAll(prefixDir, os.ModePerm|os.ModeDir)

			r1 := &result.Result{
				Student:     "Jane",
				Points:      2,
				Total:       2,
				Percentage:  100,
				SubmittedAt: time.Now(),
				Answers: []*result.Answer{
					{
						Type:    test.TaskSingle,
						Value:   "1",
						Correct: true,
					},
					{
						Type:    test.TaskDetailed,
						Value:   "detailed response",
						Correct: true,
					},
				},
			}

			r2 := &result.Result{
				Student:     "John",
				Points:      2,
				Total:       2,
				Percentage:  100,
				SubmittedAt: time.Now(),
				Answers: []*result.Answer{
					{
						Type:    test.TaskMultiple,
						Value:   "1,4,5",
						Correct: true,
					},
					{
						Type:    test.TaskOpen,
						Value:   "yes",
						Correct: true,
					},
				},
			}

			r1Bytes, _ := json.Marshal(r1)
			r2Bytes, _ := json.Marshal(r2)

			os.WriteFile(filepath.Join(prefixDir, "Jane.json"), r1Bytes, os.ModePerm)
			os.WriteFile(filepath.Join(prefixDir, "John.json"), r2Bytes, os.ModePerm)
			os.WriteFile(filepath.Join(prefixDir, "Invalid.json"), []byte("not a json"), os.ModePerm)
			os.WriteFile(filepath.Join(prefixDir, "WriteOnly.json"), r1Bytes, 0333)
			os.Mkdir(filepath.Join(prefixDir, "directory.json"), os.ModePerm|os.ModeDir)

			results, err := result.GetForTest("prefix")

			assert.NoError(t, err)
			assert.Len(t, results, 2)
			assert.Equal(t, r1.Student, results[0].Student)
			assert.Equal(t, r2.Student, results[1].Student)
		},
	)

	t.Run("should return error if prefix dir does not exist", func(t *testing.T) {
		results, err := result.GetForTest("does not exist")
		assert.Nil(t, results)
		assert.Error(t, err)
	})

	t.Run("should return error if prefix dir cannot be read", func(t *testing.T) {
		os.Mkdir(filepath.Join(tmp, "writeonly prefix"), 0333)
		results, err := result.GetForTest("writeonly prefix")
		assert.Nil(t, results)
		assert.Error(t, err)
	})
}

func TestAvailableResults(t *testing.T) {
	t.Run("should return slice of available results", func(t *testing.T) {
		tmp := t.TempDir()
		result.Init(result.Config{
			Path:      tmp,
			Overwrite: false,
		})

		iterations := 50

		for i := range iterations {
			os.Mkdir(filepath.Join(tmp, strconv.Itoa(i)), os.ModePerm|os.ModeDir)
		}

		a := result.AvailableResults()

		assert.Len(t, a, iterations)

		for i := range iterations {
			assert.Contains(t, a, strconv.Itoa(i))
		}
	})

	t.Run("should return nil (empty slice) if results dir cannot be read", func(t *testing.T) {
		tmp := t.TempDir()
		writeOnly := filepath.Join(tmp, "WriteOnly")

		os.Mkdir(writeOnly, 0333)

		result.Init(result.Config{
			Path: writeOnly,
		})

		assert.Nil(t, result.AvailableResults())
	})
}

func TestSave(t *testing.T) {
	t.Run("should save valid result", func(t *testing.T) {
		tmp := t.TempDir()
		result.Init(result.Config{
			Path: tmp,
		})

		res := &result.Result{
			Student: "John",
			Points:  100,
		}

		err := result.Save(res, "dir")

		assert.NoError(t, err)
		assert.FileExists(t, filepath.Join(tmp, "dir", "John.json"))
	})

	t.Run("should return error if testName is a whitespace-only or an empty string", func(t *testing.T) {
		tmp := t.TempDir()
		result.Init(result.Config{
			Path: tmp,
		})

		res := &result.Result{
			Student: "John",
			Points:  100,
		}

		assert.Error(t, result.Save(res, ""))
		assert.Error(t, result.Save(res, " "))
		assert.Error(t, result.Save(res, "\n"))
		assert.Error(t, result.Save(res, "\t"))
	})

	t.Run("should return error if testName contains path delimeter", func(t *testing.T) {
		tmp := t.TempDir()
		result.Init(result.Config{
			Path: tmp,
		})

		res := &result.Result{
			Student: "John",
			Points:  100,
		}

		assert.Error(t, result.Save(res, "some/another"))
	})

	t.Run("should return error if result cannot be saved", func(t *testing.T) {
		tmp := t.TempDir()
		readOnly := filepath.Join(tmp, "ReadOnly")
		os.Mkdir(readOnly, 0444)

		result.Init(result.Config{
			Path: readOnly,
		})

		res := &result.Result{
			Student: "John",
			Points:  100,
		}

		assert.Error(t, result.Save(res, "ReadOnly"))
	})

	t.Run("should return error if result cannot be saved", func(t *testing.T) {
		tmp := t.TempDir()
		file := filepath.Join(tmp, "file")
		os.WriteFile(file, []byte{}, 0444)

		result.Init(result.Config{
			Path: file,
		})

		res := &result.Result{
			Student: "John",
			Points:  100,
		}

		assert.Error(t, result.Save(res, "file"))
	})

	t.Run("should not overwrite existing result if this option is not enabled", func(t *testing.T) {
		tmp := t.TempDir()
		result.Init(result.Config{
			Path:      tmp,
			Overwrite: false, // Disable result overwriting.
		})

		result.Save(&result.Result{
			Student: "John",
			Points:  0,
		}, "dir")

		result.Save(&result.Result{
			Student: "John",
			Points:  100,
		}, "dir")

		resultFile := filepath.Join(tmp, "dir", "John.json")
		resultBytes, _ := os.ReadFile(resultFile)

		var actual result.Result

		json.Unmarshal(resultBytes, &actual)

		assert.Equal(t, 0, actual.Points)
	})

	t.Run("should overwrite existing result if this option is enabled", func(t *testing.T) {
		tmp := t.TempDir()
		result.Init(result.Config{
			Path:      tmp,
			Overwrite: true, // Enable result overwriting.
		})

		result.Save(&result.Result{
			Student: "John",
			Points:  0,
		}, "dir")

		result.Save(&result.Result{
			Student: "John",
			Points:  100,
		}, "dir")

		resultFile := filepath.Join(tmp, "dir", "John.json")
		resultBytes, _ := os.ReadFile(resultFile)

		var actual result.Result

		json.Unmarshal(resultBytes, &actual)

		assert.Equal(t, 100, actual.Points)
	})
}
