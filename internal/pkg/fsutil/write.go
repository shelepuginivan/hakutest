package fsutil

import (
	"os"
	"path/filepath"
)

// CreateAll creates file and its parent directories.
// If file exists, it is truncated.
func CreateAll(file string) (*os.File, error) {
	dir := filepath.Dir(file)

	if err := os.MkdirAll(dir, os.ModePerm|os.ModeDir); err != nil {
		return nil, err
	}

	return os.Create(file)
}

// WriteAll creates file and all its parent directories.
// It writes data to the file.
func WriteAll(file string, data []byte) error {
	dir := filepath.Dir(file)

	if err := os.MkdirAll(dir, os.ModePerm|os.ModeDir); err != nil {
		return err
	}

	return os.WriteFile(file, data, os.ModePerm)
}
