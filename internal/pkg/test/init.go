package test

import (
	"github.com/shelepuginivan/hakutest/internal/pkg/fsutil"
	"github.com/shelepuginivan/hakutest/internal/pkg/paths"
)

var testsDirectory = paths.Tests

func Init(testsDir string) {
	if fsutil.DirExists(testsDir) {
		testsDirectory = testsDir
	}
}
