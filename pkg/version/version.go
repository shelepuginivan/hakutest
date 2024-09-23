// Package version provides Hakutest version.
package version

import (
	"fmt"
	"runtime"
)

// Version of Hakutest installation.
var Version = fmt.Sprintf(
	"%s (%s/%s)",
	"Hakutest 1.1.0",
	runtime.GOARCH,
	runtime.GOOS,
)
