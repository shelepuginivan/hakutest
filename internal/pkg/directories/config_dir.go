package directories

import (
	"os"
	"path/filepath"
)

func Config() string {
	configDir, err := os.UserConfigDir()

	if err != nil {
		return "hakutest"
	}

	return filepath.Join(configDir, "hakutest")
}
