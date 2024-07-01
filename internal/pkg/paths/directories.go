package paths

import (
	"path/filepath"

	"github.com/adrg/xdg"
)

// Path to the tests directory.
var Tests = filepath.Join(xdg.DataHome, "hakutest", "tests")
