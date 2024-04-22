package results

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/shelepuginivan/hakutest/internal/pkg/application"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
	"gopkg.in/yaml.v3"
)

// ResultsService is a struct that provides methods for manipulating Result structures.
type ResultsService struct {
	app *application.App
}

// NewService returns a ResultsService instance.
func NewService(app *application.App) *ResultsService {
	return &ResultsService{app: app}
}

// CompareAnswers reports whether answers match.
// It is case-insensitive and ignores leading and trailing spaces.
func (s ResultsService) CompareAnswers(received, expected string) bool {
	return strings.TrimSpace(strings.ToLower(received)) == strings.TrimSpace(strings.ToLower(expected))
}

// CheckAnswers returns the results of the test.
func (s ResultsService) CheckAnswers(t test.Test, answers map[string][]string) TestResults {
	submittedAt := time.Now()
	student := strings.Join(answers["student"], "")

	results := TestResults{
		Student:     student,
		SubmittedAt: submittedAt,
		Test: TestInfo{
			Title:  t.Title,
			Author: t.Author,
			Sha256: t.Sha256Sum(),
		},
		Results: Results{
			Points:     0,
			Total:      len(t.Tasks),
			Percentage: 0,
			Tasks:      make(map[string]TaskResult),
		},
	}

	for i, answer := range answers {
		index, err := strconv.Atoi(i)

		if err != nil {
			continue
		}

		studentAnswer := strings.Join(answer, ",")
		correctAnswer := t.Tasks[index].Answer
		isCorrect := s.CompareAnswers(studentAnswer, correctAnswer)

		results.Results.Tasks[strconv.Itoa(index+1)] = TaskResult{
			Answer:  studentAnswer,
			Correct: isCorrect,
		}

		if isCorrect {
			results.Results.Points++
		}
	}

	results.Results.Percentage = 100 * results.Results.Points / len(t.Tasks)

	return results
}

// CheckAnswersWithFiles returns the results of the test, but also saves uploaded files.
func (s ResultsService) CheckAnswersWithFiles(testName string, t test.Test, answers map[string][]string, files map[string][]*multipart.FileHeader) TestResults {
	results := s.CheckAnswers(t, answers)

	for i, fileHeaders := range files {
		index, err := strconv.Atoi(i)
		if err != nil {
			continue
		}

		for fileIndex, file := range fileHeaders {
			src, err := file.Open()
			if err != nil {
				continue
			}
			defer src.Close()

			filename := fmt.Sprintf("%s-%d%s", i, fileIndex+1, filepath.Ext(file.Filename))

			if err := s.SaveSubmittedFile(src, testName, results.Student, filename); err != nil {
				continue
			}

			results.Results.Tasks[strconv.Itoa(index+1)] = TaskResult{
				Answer:  filename,
				Correct: true,
			}
		}

		results.Results.Points += 1
		results.Results.Percentage = 100 * results.Results.Points / len(t.Tasks)
	}

	return results
}

// GetResultsList retrieves a list of results names from the results directory specified in the configuration.
func (s ResultsService) GetResultsList() []string {
	resultsList := []string{}

	entries, err := os.ReadDir(s.app.Config.General.ResultsDirectory)

	if err != nil {
		return resultsList
	}

	for _, file := range entries {
		resultsName := file.Name()

		if file.IsDir() {
			resultsList = append(resultsList, resultsName)
		}
	}

	return resultsList
}

// GetTestResultsDirectory returns the absolute path of the test results directory by its name.
// It doesn't check whether a test with this name or a results directory associated with it exists.
func (s ResultsService) GetTestResultsDirectory(testName string) string {
	name := strings.TrimSuffix(testName, ".json")

	return filepath.Join(s.app.Config.General.ResultsDirectory, name)
}

// GetResultsOfTest retrieves all results of the test from the results directory specified in the configuration.
// The name is the filename of the test.
func (s ResultsService) GetResultsOfTest(testName string) ([]TestResults, error) {
	results := []TestResults{}
	testResultsDir := s.GetTestResultsDirectory(testName)
	entries, err := os.ReadDir(testResultsDir)

	if err != nil {
		return nil, err
	}

	for _, file := range entries {
		if file.IsDir() {
			continue
		}

		data, err := os.ReadFile(filepath.Join(testResultsDir, file.Name()))

		if err != nil {
			continue
		}

		entry := TestResults{}

		if yaml.Unmarshal(data, &entry) != nil {
			continue
		}

		results = append(results, entry)
	}

	return results, nil
}

// Save saves test results in the results directory specified in the configuration.
// The results are saved in a subdirectory name.
func (s ResultsService) Save(r TestResults, testName string) error {
	testResultsDirectory := s.GetTestResultsDirectory(testName)
	resultsFilePath := filepath.Join(testResultsDirectory, r.Student+".txt")

	_, err := os.Stat(resultsFilePath)
	if !os.IsNotExist(err) && !s.app.Config.General.OverwriteResults {
		// Test was already submitted by this student
		return err
	}

	err = os.MkdirAll(testResultsDirectory, os.ModeDir|os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return err
	}

	data, err := yaml.Marshal(r)
	if err != nil {
		return err
	}

	return os.WriteFile(resultsFilePath, data, 0666)
}

// SaveSubmittedFile saves file attached to the test solution in the respective results directory.
func (s ResultsService) SaveSubmittedFile(file multipart.File, testName, student, filename string) error {
	testResultsDirectory := s.GetTestResultsDirectory(testName)

	studentSubmittedFilesDir := filepath.Join(testResultsDirectory, strings.ReplaceAll(student, "/", ""))

	err := os.MkdirAll(studentSubmittedFilesDir, os.ModeDir|os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return err
	}

	fullpath := filepath.Join(studentSubmittedFilesDir, filename)

	_, err = os.Stat(fullpath)
	if !os.IsNotExist(err) && !s.app.Config.General.OverwriteResults {
		// Submitted file was already written
		return err
	}

	dst, err := os.Create(fullpath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	return err
}

// Remove removes the directory that stores test solutions.
// It uses name of the test (i.e. the name of the associated directory).
func (s ResultsService) Remove(testName string) error {
	return os.RemoveAll(s.GetTestResultsDirectory(testName))
}
