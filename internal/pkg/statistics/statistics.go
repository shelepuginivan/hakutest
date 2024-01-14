package statistics

import (
	"github.com/shelepuginivan/hakutest/internal/pkg/results"
)

type Statistics struct {
	Entries []results.TestResults
}

func New(entries []results.TestResults) Statistics {
	return Statistics{Entries: entries}
}
