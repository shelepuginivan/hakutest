package directories

import (
	"os"
	"path/filepath"
)

func Data() string {
	cacheDir, err := os.UserCacheDir()

	if err != nil {
		return "data"
	}

	return filepath.Join(cacheDir, "hakutest")
}
