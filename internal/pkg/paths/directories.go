package paths

import (
	"path/filepath"

	"github.com/adrg/xdg"
)

// Patht to the results directory.
var Results = filepath.Join(xdg.DataHome, "hakutest", "results")

// Path to the tests directory.
var Tests = filepath.Join(xdg.DataHome, "hakutest", "tests")
