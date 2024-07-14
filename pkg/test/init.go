package test

import (
	"sync"

	"github.com/shelepuginivan/hakutest/internal/pkg/fsutil"
	"github.com/shelepuginivan/hakutest/internal/pkg/paths"
)

// Test configuration section.
type Config struct {
	// Directory where tests are stored.
	Path string `yaml:"path"`
}

var (
	testsDirectory = paths.Tests

	mu sync.Mutex
)

func Init(cfg Config) {
	mu.Lock()
	defer mu.Unlock()

	if fsutil.DirExists(cfg.Path) {
		testsDirectory = cfg.Path
	}
}
