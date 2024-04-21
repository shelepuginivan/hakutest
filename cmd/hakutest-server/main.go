package main

import (
	"context"
	"net/http"
	"os"
	"path/filepath"

	"github.com/getlantern/systray"
	"github.com/shelepuginivan/hakutest/internal/app/server"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/directories"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
)

const (
	appName  = "Hakutest Server"
	iconPath = "web/static/img/favicon.ico"
)

func main() {
	systray.Run(onReady, func() {})
}

func getIcon(s string) []byte {
	b, err := os.ReadFile(s)
	if err != nil {
		panic(err)
	}
	return b
}

func onReady() {
	serverI18n := i18n.New().Server

	systray.SetTitle(appName)
	systray.SetTooltip(appName)
	systray.SetIcon(getIcon(filepath.Join(directories.Executable(), iconPath)))

	mStop := systray.AddMenuItem(serverI18n.StopTitle, serverI18n.StopTooltip)

	port := config.New().Server.Port
	r := server.NewEngine(test.NewService(), results.NewService())
	srv := server.NewServer(r, port)

	go func() {
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			panic(err)
		}
	}()

	go func() {
		<-mStop.ClickedCh
		systray.Quit()

		err := srv.Shutdown(context.Background())
		if err != nil {
			panic(err)
		}
	}()
}
