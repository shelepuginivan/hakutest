package runtime

import (
	"os"
	"path/filepath"
)

func ExecutableDir() string {
	executable, err := os.Executable()

	if err != nil {
		return "."
	}

	return filepath.Dir(executable)
}
