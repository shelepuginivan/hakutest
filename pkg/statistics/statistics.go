package statistics

import "github.com/shelepuginivan/hakutest/pkg/result"

type Statistics struct {
	Name    string
	Total   int
	Results []*result.Result
}

func New(testName string, results []*result.Result) *Statistics {
	s := &Statistics{
		Name:    testName,
		Results: results,
	}

	if len(results) > 0 {
		s.Total = s.Results[0].Total
	}

	return s
}

func NewFromName(testName string) (*Statistics, error) {
	results, err := result.GetForTest(testName)
	if err != nil {
		return nil, err
	}

	s := &Statistics{
		Name:    testName,
		Results: results,
	}

	if len(results) > 0 {
		s.Total = s.Results[0].Total
	}

	return s, nil
}
