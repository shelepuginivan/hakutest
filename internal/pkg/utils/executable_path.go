package utils

import (
	"os"
	"path/filepath"
)

func GetExecutablePath() string {
	executable, err := os.Executable()

	if err != nil {
		return "."
	}

	return filepath.Dir(executable)
}
