package statistics

import (
	"encoding/json"
	"io"
)

const (
	// Human readable description of JSON export format.
	DescriptionJSON = ".json (JavaScript Object Notation)"

	// JSON statistics export format.
	FormatJSON = "json"

	// JSON file MIME type.
	MimeJSON = "application/json"
)

// ToJSON returns statistics bytes in JSON format.
func (s *Statistics) ToJSON() ([]byte, error) {
	return json.Marshal(s.Results)
}

// WriteJSON writes statistics in JSON format to w.
func (s *Statistics) WriteJSON(w io.Writer) error {
	data, err := json.Marshal(s.Results)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}
