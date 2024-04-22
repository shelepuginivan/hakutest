// Package application provides app configuration layer.
package application

import (
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
)

// App represents app configuration layer.
type App struct {
	Config *config.Config // Config of the app.
	I18n   *i18n.I18n     // Internationalization of the app.
}

// New returns a new instance of App.
func New() *App {
	return &App{
		Config: config.New(),
		I18n:   i18n.New(),
	}
}
