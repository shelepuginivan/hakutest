//go:build linux || darwin

package icon

import _ "embed"

//go:embed icon_unix.png
var Icon []byte
