package fsutil_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/shelepuginivan/hakutest/internal/pkg/fsutil"
	"github.com/stretchr/testify/assert"
)

func TestDirExists(t *testing.T) {
	t.Run("should return true if directory exists", func(t *testing.T) {
		assert.True(t, fsutil.DirExists(t.TempDir()))
	})

	t.Run("should return false if directory does not exist", func(t *testing.T) {
		assert.False(t, fsutil.DirExists("this dir does not exist"))
	})

	t.Run("should return false if file exists, but is not a directory", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "file")
		os.WriteFile(path, []byte("file"), os.ModePerm)
		assert.False(t, fsutil.DirExists(path))
	})
}

func TestFileExists(t *testing.T) {
	t.Run("should return true if file exists", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "file")
		os.WriteFile(path, []byte("file"), os.ModePerm)
		assert.True(t, fsutil.FileExists(path))
	})

	t.Run("should return false if file does not exist", func(t *testing.T) {
		assert.False(t, fsutil.FileExists("this file does not exist"))
	})

	t.Run("should return false if file exists, but is not a file", func(t *testing.T) {
		assert.False(t, fsutil.FileExists(t.TempDir()))
	})
}

func TestReadIfExists(t *testing.T) {
	t.Run("should read first file that exists", func(t *testing.T) {
		expected := []byte("content of the file")
		path := filepath.Join(t.TempDir(), "file")
		os.WriteFile(path, expected, os.ModePerm)

		actual, err := fsutil.ReadIfExists("1", "2", "3", "4", path)
		assert.NoError(t, err)
		assert.True(t, bytes.Equal(expected, actual))
	})

	t.Run("should read first file that exists (multiple files exist)", func(t *testing.T) {
		expected := []byte("content of the file")
		path := filepath.Join(t.TempDir(), "file")
		os.WriteFile(path, expected, os.ModePerm)

		anotherPath := filepath.Join(t.TempDir(), "file")
		os.WriteFile(anotherPath, []byte("another file"), os.ModePerm)

		actual, err := fsutil.ReadIfExists("1", path, anotherPath)
		assert.NoError(t, err)
		assert.True(t, bytes.Equal(expected, actual))
	})

	t.Run("should return error if neither file exist", func(t *testing.T) {
		data, err := fsutil.ReadIfExists("1", "2", "3", "4", "5")
		assert.Nil(t, data)
		assert.Error(t, err)
	})
}

func TestCreateAll(t *testing.T) {
	t.Run("should create new file and parent directories", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "1", "2", "3", "file")
		file, err := fsutil.CreateAll(path)
		assert.NoError(t, err)
		file.Close()
	})

	t.Run("should return error if file cannot be created", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "1")
		filePath := filepath.Join(path, "2", "file")
		os.WriteFile(path, []byte("exists"), os.ModePerm)
		_, err := fsutil.CreateAll(filePath)
		assert.Error(t, err)
	})
}

func TestWriteAll(t *testing.T) {
	t.Run("should write file and create all parent directories", func(t *testing.T) {
		expected := []byte("content")
		path := filepath.Join(t.TempDir(), "1", "2", "3", "file")

		err := fsutil.WriteAll(path, expected)
		assert.NoError(t, err)

		actual, err := os.ReadFile(path)
		assert.NoError(t, err)
		assert.True(t, bytes.Equal(expected, actual))
	})

	t.Run("should return error if file cannot be written", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "1")
		filePath := filepath.Join(path, "2", "file")
		os.WriteFile(path, []byte("exists"), os.ModePerm)
		err := fsutil.WriteAll(filePath, []byte{})
		assert.Error(t, err)
	})
}
