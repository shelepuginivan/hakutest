package runtime

import (
	"os"
	"path/filepath"
)

func ConfigDir() string {
	configDir, err := os.UserConfigDir()

	if err != nil {
		return "hakutest"
	}

	return filepath.Join(configDir, "hakutest")
}
