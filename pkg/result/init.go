package result

import (
	"sync"

	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/paths"
)

var (
	overwriteResults = false
	resultsDirectory = paths.Results

	mu sync.Mutex
)

func Init(cfg *config.Config) {
	mu.Lock()
	defer mu.Unlock()

	overwriteResults = cfg.OverwriteResults
	resultsDirectory = cfg.ResultsDirectory
}
