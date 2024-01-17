package directories

import (
	"os"
	"path/filepath"
)

func Executable() string {
	executable, err := os.Executable()

	if err != nil {
		return "."
	}

	return filepath.Dir(executable)
}
