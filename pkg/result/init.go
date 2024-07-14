package result

import (
	"sync"

	"github.com/shelepuginivan/hakutest/internal/pkg/paths"
)

// Result configuration section.
type Config struct {
	// Whether to overwrite results on resend.
	Overwrite bool `yaml:"overwrite"`

	// Directory where results are stored.
	Path string `yaml:"path"`

	// Whether to show results on submission.
	Show bool `yaml:"show"`
}

var (
	overwriteResults = false
	resultsDirectory = paths.Results

	mu sync.Mutex
)

func Init(cfg Config) {
	mu.Lock()
	defer mu.Unlock()

	overwriteResults = cfg.Overwrite
	resultsDirectory = cfg.Path
}
