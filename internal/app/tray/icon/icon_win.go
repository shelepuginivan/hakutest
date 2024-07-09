//go:build windows

package icon

import _ "embed"

//go:embed icon_win.ico
var Icon []byte
