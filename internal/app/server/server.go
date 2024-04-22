package server

import (
	"fmt"
	"net/http"

	"github.com/shelepuginivan/hakutest/internal/pkg/application"
)

func NewServer(handler http.Handler, port int) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}
}

func New(app *application.App) *http.Server {
	engine := NewEngine(app)
	addr := fmt.Sprintf(":%d", app.Config.Server.Port)

	return &http.Server{
		Addr:    addr,
		Handler: engine,
	}
}
