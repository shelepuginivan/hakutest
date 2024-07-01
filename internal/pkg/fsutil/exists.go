// Package fsutil provides file system utilities.
package fsutil

import (
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
