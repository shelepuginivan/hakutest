package statistics

import "github.com/shelepuginivan/hakutest/pkg/result"

type Statistics struct {
	Name    string
	Results []*result.Result
}

func New(testName string, results []*result.Result) *Statistics {
	return &Statistics{
		Name:    testName,
		Results: results,
	}
}

func NewFromName(testName string) (*Statistics, error) {
	results, err := result.GetForTest(testName)
	if err != nil {
		return nil, err
	}

	return &Statistics{
		Name:    testName,
		Results: results,
	}, nil
}
