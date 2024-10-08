// Package web provides embedded file systems for the web server.
package web

import "embed"

// Templates directory.
//
//go:embed partials
//go:embed templates
var Templates embed.FS

// Static files directories.
//
//go:embed css
//go:embed fonts
//go:embed img
//go:embed js
//go:embed vendor
var Static embed.FS
