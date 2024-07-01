// Package paths provides common paths for the application.
package paths

import (
	"path/filepath"

	"github.com/adrg/xdg"
)

// Path to the configuration file.
var Config = filepath.Join(xdg.ConfigHome, "hakutest", "config.yaml")
