// Package version provides Hakutest version.
package version

import (
	_ "embed"

	"fmt"
	"runtime"
)

//go:embed VERSION
var version string

// Version of Hakutest installation.
var Version = fmt.Sprintf(
	"Hakutest %s (%s/%s)",
	version,
	runtime.GOARCH,
	runtime.GOOS,
)
