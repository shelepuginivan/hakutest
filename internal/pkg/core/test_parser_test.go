package core

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetTestPath(t *testing.T) {
	cases := []struct {
		name   string
		suffix string
	}{
		{name: "test", suffix: "/test.json"},
		{name: "Name of the test", suffix: "/Name of the test.json"},
		{name: "1", suffix: "/1.json"},
		{name: "1.json", suffix: "/1.json"},
		{name: "New test.json", suffix: "/New test.json"},
	}

	for _, c := range cases {
		assert.True(t, strings.HasSuffix(GetTestPath(c.name), c.suffix))
	}
}

func TestSha256Sum(t *testing.T) {
	tests := []Test{
		{
			Title:       "A test",
			Description: "Description of the test",
			Author:      "John Doe",
			Subject:     "test",
			Target:      "test",
			Institution: "test",
			CreatedAt:   time.Now(),
			ExpiresIn:   time.Now(),
			Tasks: []Task{
				{
					Type:    "single",
					Text:    "task 1",
					Options: []string{"1", "2", "3"},
					Answer:  "1",
				},
			},
		},
		{
			Title:       "1",
			Description: "2",
			Author:      "3",
			Subject:     "4",
			Target:      "5",
			Institution: "6",
			CreatedAt:   time.Now(),
			ExpiresIn:   time.Now(),
			Tasks: []Task{
				{
					Type:    "multiple",
					Text:    "7",
					Options: []string{"8", "9", "10"},
					Answer:  "1,2",
					Attachment: Attachment{
						Type: "image",
						Name: "example",
						Src:  "https://example.com/img.png",
					},
				},
			},
		},
	}

	for _, test := range tests {
		hasher := sha256.New()
		data, err := json.Marshal(test)

		if err != nil {
			t.Logf(fmt.Sprintf("Failed to marshal test: %s", err))
		}

		hasher.Write(data)

		expected := hex.EncodeToString(hasher.Sum(nil))
		actual := test.Sha256Sum()

		assert.Equal(t, actual, expected)
	}
}
