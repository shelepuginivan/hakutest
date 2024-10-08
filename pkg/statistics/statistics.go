// Package statistics provides methods for exporting statistics of the testing
// results.
package statistics

import "github.com/shelepuginivan/hakutest/pkg/result"

// Statistics is a struct that encapsulates results.
// Results can be exported to various data formats.
type Statistics struct {
	// Name of the statistics report, typically the same as the test name.
	Name string

	// Total number of tasks and points.
	Total   int
	Results []*result.Result
}

// New returns a new instance of Statistics.
func New(testName string, results []*result.Result) *Statistics {
	s := &Statistics{
		Name:    testName,
		Results: results,
	}

	// Set s.Total to the maximal value of total tasks among results. This is
	// required since tests can be modified, hence total number of tasks of
	// results may differ.
	for _, r := range s.Results {
		if r.Total > s.Total {
			s.Total = r.Total
		}
	}

	return s
}

// NewFromSaved returns a new instance of Statistics.
// It reads results from the results directory.
// E.g. if the testName is `foo`, it reads `/results/directory/foo`.
func NewFromSaved(testName string) (*Statistics, error) {
	results, err := result.GetForTest(testName)
	if err != nil {
		return nil, err
	}

	return New(testName, results), nil
}
