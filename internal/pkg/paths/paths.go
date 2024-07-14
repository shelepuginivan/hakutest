// Package paths provides XDG-compliant common paths for the application.
//
// This package creates necessary directories such as cache, config, and data
// directories if they don't exist.
//
// [XDG Base Directory]: https://wiki.archlinux.org/title/XDG_Base_Directory
package paths

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

func init() {
	err := os.MkdirAll(filepath.Join(xdg.CacheHome, "hakutest"), os.ModePerm|os.ModeDir)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(filepath.Join(xdg.ConfigHome, "hakutest"), os.ModePerm|os.ModeDir)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(filepath.Join(xdg.DataHome, "hakutest"), os.ModePerm|os.ModeDir)
	if err != nil {
		panic(err)
	}
}

// Files.
var (
	// Path to the configuration file.
	Config = filepath.Join(xdg.ConfigHome, "hakutest", "config.yaml")

	// Path to the log file.
	Logs = filepath.Join(xdg.CacheHome, "hakutest", "hakutest.log")

	// Path to the default SQLite database with encrypted user data.
	UserDB = filepath.Join(xdg.CacheHome, "hakutest", "users.db")
)

// Directories.
var (
	// Path to the results directory.
	Results = filepath.Join(xdg.DataHome, "hakutest", "results")

	// Path to the tests directory.
	Tests = filepath.Join(xdg.DataHome, "hakutest", "tests")
)
