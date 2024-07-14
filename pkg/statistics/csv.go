package statistics

import (
	"bytes"
	"encoding/csv"
	"io"
	"strconv"
	"time"
)

const (
	// Human readable description of CSV export format.
	DescriptionCSV = ".csv (Comma-separated values)"

	// CSV statistics export format.
	FormatCSV = "csv"

	// CSV file MIME type.
	MimeCSV = "text/csv"
)

// ToCSV returns statistics bytes in CSV format.
// It calls WriteCSV method of s internally.
func (s *Statistics) ToCSV() ([]byte, error) {
	b := bytes.NewBuffer(nil)
	if err := s.WriteCSV(b); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// WriteCSV writes statistics in CSV format to w.
func (s *Statistics) WriteCSV(w io.Writer) (err error) {
	csvWriter := csv.NewWriter(w)

	for _, r := range s.Results {
		record := []string{
			r.Student,
			strconv.Itoa(r.Points),
			strconv.Itoa(r.Percentage),
			r.SubmittedAt.Format(time.DateTime),
		}

		for _, a := range r.Answers {
			record = append(record, a.Value)
		}

		if err = csvWriter.Write(record); err != nil {
			return err
		}
	}

	csvWriter.Flush()
	return csvWriter.Error()
}
