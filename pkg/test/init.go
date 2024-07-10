package test

import (
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/fsutil"
	"github.com/shelepuginivan/hakutest/internal/pkg/paths"
)

var testsDirectory = paths.Tests

func Init(cfg *config.Config) {
	if fsutil.DirExists(cfg.TestsDirectory) {
		testsDirectory = cfg.TestsDirectory
	}
}
