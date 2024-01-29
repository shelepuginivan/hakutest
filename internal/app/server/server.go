package server

import (
	"fmt"
	"net/http"
)

func NewServer(handler http.Handler, port int) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}
}
