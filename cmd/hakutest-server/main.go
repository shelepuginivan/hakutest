package main

import (
	"context"
	_ "embed"
	"net/http"

	"github.com/getlantern/systray"
	"github.com/shelepuginivan/hakutest/internal/app/server"
	"github.com/shelepuginivan/hakutest/internal/pkg/application"
)

const appName = "Hakutest Server"

//go:embed icon.ico
var icon []byte

// onReady returns a callback for the systray.Run.
func onReady(app *application.App) func() {
	return func() {
		systray.SetTitle(appName)
		systray.SetTooltip(appName)
		systray.SetIcon(icon)

		mStop := systray.AddMenuItem(
			app.I18n.Server.StopTitle,
			app.I18n.Server.StopTooltip,
		)

		srv := server.New(app)

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
}

func main() {
	systray.Run(onReady(application.New()), func() {})
}
