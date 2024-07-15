// Package fsutil provides reusable file system utility methods.
package fsutil

import (
	"fmt"
	"os"
)

// DirExists reports whether path exists and is a directory.
func DirExists(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}

	return stat.IsDir()
}

// FileExists reports whether path exists and is a file.
func FileExists(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}

	return !stat.IsDir()
}

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
