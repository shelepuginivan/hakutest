// Package web provides embedded file systems for the web server.
package web

import "embed"

//go:embed templates
var Templates embed.FS

//go:embed css fonts js img
var Static embed.FS
