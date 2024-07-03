// Package embedded provides embedded files and assets that are used multiple
// times.
package embedded

import _ "embed"

//go:embed icon.ico
var Icon []byte
