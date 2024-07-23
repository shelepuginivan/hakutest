package statistics_test

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/shelepuginivan/hakutest/pkg/result"
	"github.com/shelepuginivan/hakutest/pkg/statistics"
	"github.com/shelepuginivan/hakutest/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/xuri/excelize/v2"
)

var results = []*result.Result{
	{
		Student:     "A",
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
				Type:    test.TaskOpen,
				Value:   "open task",
				Correct: true,
			},
		},
	},
	{
		Student:     "B",
		Points:      1,
		Total:       2,
		Percentage:  50,
		SubmittedAt: time.Now(),
		Answers: []*result.Answer{
			{
				Type:    test.TaskSingle,
				Value:   "2",
				Correct: false,
			},
			{
				Type:    test.TaskOpen,
				Value:   "open task",
				Correct: true,
			},
		},
	},
}

func TestNew(t *testing.T) {
	t.Run("should return new instance of statistics (field Total is equal)", func(t *testing.T) {
		s := statistics.New("test", results)

		assert.Equal(t, "test", s.Name)
		assert.Equal(t, 2, s.Total)
		assert.True(t, reflect.DeepEqual(results, s.Results))
	})

	t.Run("should return new instance of statistics", func(t *testing.T) {
		res := []*result.Result{
			results[0],
			results[1],
			{
				Student:     "C",
				Total:       3,
				Percentage:  100,
				SubmittedAt: time.Now(),
				Answers: []*result.Answer{
					{
						Type:    test.TaskSingle,
						Value:   "1",
						Correct: true,
					},
					{
						Type:    test.TaskOpen,
						Value:   "open task",
						Correct: true,
					},
					{
						Type:    test.TaskDetailed,
						Value:   "my detailed answer",
						Correct: true,
					},
				},
			},
			{
				Student:    "D",
				Points:     0,
				Total:      1,
				Percentage: 0,
				Answers: []*result.Answer{
					{
						Type:    test.TaskDetailed,
						Value:   "",
						Correct: false,
					},
				},
			},
		}

		s := statistics.New("test", res)

		assert.Equal(t, "test", s.Name)
		assert.Equal(t, 3, s.Total)
		assert.True(t, reflect.DeepEqual(res, s.Results))

	})
}

func TestNewFromSaved(t *testing.T) {
	tmp := t.TempDir()
	result.Init(result.Config{
		Path: tmp,
	})

	results := []*result.Result{
		{Student: "A", Total: 2},
		{Student: "B", Total: 2},
		{Student: "C", Total: 5},
		{Student: "D", Total: 2},
		{Student: "E", Total: 3},
		{Student: "F", Total: 2},
		{Student: "G", Total: 1},
		{Student: "H", Total: 2},
	}

	for _, r := range results {
		err := result.Save(r, "test")
		if err != nil {
			panic(err)
		}
	}

	s, err := statistics.NewFromSaved("test")

	assert.NoError(t, err)
	assert.Equal(t, "test", s.Name)
	assert.Equal(t, 5, s.Total)
	assert.True(t, reflect.DeepEqual(results, s.Results))
}

func TestWriteCSV(t *testing.T) {
	b := bytes.NewBuffer(nil)
	s := statistics.New("test", results)

	err := s.WriteCSV(b)
	assert.NoError(t, err)

	r := csv.NewReader(b)

	entries, _ := r.ReadAll()

	for row, entry := range entries {
		assert.Equal(t, s.Results[row].Student, entry[0])
		assert.Equal(t, strconv.Itoa(s.Results[row].Points), entry[1])
		assert.Equal(t, strconv.Itoa(s.Results[row].Percentage), entry[2])
		assert.Equal(t, s.Results[row].SubmittedAt.Format(time.DateTime), entry[3])

		for col, ans := range entry[4:] {
			assert.Equal(t, s.Results[row].Answers[col].Value, ans)
		}
	}
}

func TestToCSV(t *testing.T) {
	s := statistics.New("test", results)
	data, err := s.ToCSV()
	assert.NoError(t, err)
	b := bytes.NewBuffer(data)

	r := csv.NewReader(b)

	entries, _ := r.ReadAll()

	for row, entry := range entries {
		assert.Equal(t, s.Results[row].Student, entry[0])
		assert.Equal(t, strconv.Itoa(s.Results[row].Points), entry[1])
		assert.Equal(t, strconv.Itoa(s.Results[row].Percentage), entry[2])
		assert.Equal(t, s.Results[row].SubmittedAt.Format(time.DateTime), entry[3])

		for col, ans := range entry[4:] {
			assert.Equal(t, s.Results[row].Answers[col].Value, ans)
		}
	}
}

func TestWriteJSON(t *testing.T) {
	b := bytes.NewBuffer(nil)
	s := statistics.New("test", results)

	err := s.WriteJSON(b)
	assert.NoError(t, err)

	expected, _ := json.Marshal(s.Results)
	actual := b.Bytes()

	assert.True(t, bytes.Equal(expected, actual))
}

func TestToJSON(t *testing.T) {
	s := statistics.New("test", results)

	expected, _ := json.Marshal(s.Results)
	actual, err := s.ToJSON()

	assert.NoError(t, err)
	assert.True(t, bytes.Equal(expected, actual))
}

func TestWriteXLSX(t *testing.T) {
	b := bytes.NewBuffer(nil)
	s := statistics.New("test", results)

	err := s.WriteXLSX(b)
	assert.NoError(t, err)

	file, _ := excelize.OpenReader(b)
	sheet := "Hakutest"

	assert.Equal(t, 1, file.SheetCount)

	for row, res := range s.Results {
		exRow := row + 2 // 1-indexed, 1st row is a header row.
		exStudent, _ := file.GetCellValue(sheet, fmt.Sprintf("A%d", exRow))
		exPoints, _ := file.GetCellValue(sheet, fmt.Sprintf("B%d", exRow))
		exPercentage, _ := file.GetCellValue(sheet, fmt.Sprintf("C%d", exRow))
		exSubmittedAt, _ := file.GetCellValue(sheet, fmt.Sprintf("D%d", exRow))

		assert.Equal(t, res.Student, exStudent)
		assert.Equal(t, strconv.Itoa(res.Points), exPoints)
		assert.Equal(t, strconv.Itoa(res.Percentage), exPercentage)
		assert.Equal(t, res.SubmittedAt.Format("1/2/06 15:04"), exSubmittedAt)

		for col, ans := range res.Answers {
			// 1-indexed, 1st, 2nd, 3rd, and 4th are for student, points,
			// percentage, and submission time respectively.
			ansCell, _ := excelize.CoordinatesToCellName(col+5, exRow)
			exAnswer, _ := file.GetCellValue(sheet, ansCell)

			assert.Equal(t, ans.HumanReadable(), exAnswer)
		}
	}
}
