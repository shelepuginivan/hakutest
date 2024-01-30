package statistics

import (
	"fmt"

	"github.com/shelepuginivan/hakutest/internal/pkg/results"
)

type StatisticsService struct {
	r *results.ResultsService
}

func NewService(r *results.ResultsService) *StatisticsService {
	return &StatisticsService{r: r}
}

func (s StatisticsService) Export(testName, dest, format string) error {
	res, err := s.r.GetResultsOfTest(testName)

	if err != nil {
		return err
	}

	stats := New(res)

	switch format {
	case FormatExcel:
		return stats.ExportToExcel(dest)
	case FormatImage:
		return stats.ExportToPng(dest)
	case FormatTable:
		stats.ExportToTable().Print()
		return nil
	default:
		return fmt.Errorf("unknown format %s", format)
	}
}
