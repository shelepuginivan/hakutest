package fsutil

import (
	"fmt"
	"os"
)

// ReadIfExists checks files until existing one is found.
// It reads the first existing file and returns read operation result.
// If neither of provided files exist, error is returned.
func ReadIfExists(files ...string) ([]byte, error) {
	for _, file := range files {
		if !FileExists(file) {
			continue
		}

		return os.ReadFile(file)
	}

	return nil, fmt.Errorf("neither of files exist")
}
