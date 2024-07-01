// Package web provides embedded file systems for the web server.
package web

import "embed"

// Templates directory.
//
//go:embed templates
var Templates embed.FS

// Static files directories.
var Static embed.FS
