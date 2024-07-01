package config

import (
	"path/filepath"

	"github.com/adrg/xdg"
)

// configFile returns path to the config file.
func configFile() string {
	return filepath.Join(xdg.ConfigHome, "hakutest", "config.yaml")
}
