package test

import (
	"sync"

	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/fsutil"
	"github.com/shelepuginivan/hakutest/internal/pkg/paths"
)

var (
	testsDirectory = paths.Tests

	mu sync.Mutex
)

func Init(cfg *config.Config) {
	mu.Lock()
	defer mu.Unlock()

	if fsutil.DirExists(cfg.TestsDirectory) {
		testsDirectory = cfg.TestsDirectory
	}
}
