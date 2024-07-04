package paths

import (
	"path/filepath"

	"github.com/adrg/xdg"
)

// Path to the directory with encrypted user data.
var Users = filepath.Join(xdg.CacheHome, "hakutest", "users")

// Path to the results directory.
var Results = filepath.Join(xdg.DataHome, "hakutest", "results")

// Path to the tests directory.
var Tests = filepath.Join(xdg.DataHome, "hakutest", "tests")
