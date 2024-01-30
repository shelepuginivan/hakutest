package main

import (
	"context"
	"net/http"
	"os"

	"github.com/getlantern/systray"
	"github.com/shelepuginivan/hakutest/internal/app/server"
	"github.com/shelepuginivan/hakutest/internal/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
)

const IconPath = "web/static/img/favicon.ico"

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

	systray.SetTitle("Hakutest server")
	systray.SetTooltip("Hakutest")
	systray.SetIcon(getIcon(IconPath))

	mStop := systray.AddMenuItem(serverI18n.StopTitle, serverI18n.StopTooltip)

	port := config.New().Server.Port
	r := server.NewRouter(test.NewService(), results.NewService())
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
