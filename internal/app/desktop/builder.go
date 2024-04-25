package desktop

import (
	"github.com/shelepuginivan/hakutest/internal/pkg/application"
)

type Builder struct {
	app *application.App
}

func NewBuilder(app *application.App) *Builder {
	return &Builder{app: app}
}
