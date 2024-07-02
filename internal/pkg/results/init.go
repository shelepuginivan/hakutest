package results

import (
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/paths"
)

var overwriteResults = false
var resultsDirectory = paths.Results

func Init(cfg *config.Config) {
	overwriteResults = cfg.OverwriteResults
	resultsDirectory = cfg.ResultsDirectory
}
