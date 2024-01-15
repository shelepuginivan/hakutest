package runtime

import (
	"os"
	"path/filepath"
)

func DataDir() string {
	cacheDir, err := os.UserCacheDir()

	if err != nil {
		return "data"
	}

	return filepath.Join(cacheDir, "hakutest")
}
