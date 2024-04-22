package statistics

import (
	"fmt"

	"github.com/shelepuginivan/hakutest/internal/pkg/results"
)

// StatisticsService is a struct that provides methods for manipulating Statistics structures.
type StatisticsService struct {
	app *application.App
	r   *results.ResultsService
}

// NewService returns a StatisticsService instance.
func NewService(app *application.App, r *results.ResultsService) *StatisticsService {
	return &StatisticsService{r: r}
}

// Export retrieves statistics of the test and exports in to a specified format.
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
