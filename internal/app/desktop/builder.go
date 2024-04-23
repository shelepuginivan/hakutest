package desktop

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/hakutest/internal/pkg/application"
)

type Builder struct {
	app *application.App
	win *gtk.Window
}

func NewBuilder(app *application.App, win *gtk.Window) *Builder {
	return &Builder{
		app: app,
		win: win,
	}
}
