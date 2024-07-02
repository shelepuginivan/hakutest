package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/getlantern/systray"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/shelepuginivan/hakutest/internal/pkg/server"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
	"github.com/shelepuginivan/hakutest/internal/pkg/trayutil"
)

var (
	cfg *config.Config
	srv *http.Server
)

func init() {
	cfg = config.New()

	flag.BoolVar(&cfg.Debug, "debug", cfg.Debug, "Run in debug mode")
	flag.BoolVar(&cfg.Headless, "headless", cfg.Headless, "Run in headless mode (without systray icon)")
	flag.StringVar(&cfg.Lang, "lang", cfg.Lang, "Language of the interface")
	flag.IntVar(&cfg.Port, "port", cfg.Port, "Port on which server is started")
	flag.StringVar(&cfg.TestsDirectory, "tests-directory", cfg.TestsDirectory, "Directory where the test files are stored")

	flag.Parse()

	i18n.Init(cfg.Lang)
	results.Init(cfg)
	test.Init(cfg.TestsDirectory)
	srv = server.New(cfg)
}

func onReady() {
	systray.SetIcon(trayutil.Icon)
	systray.SetTitle(trayutil.Title)
	systray.SetTooltip(trayutil.Tooltip)

	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit application")

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			systray.Quit()
		}
	}()

	go func() {
		<-mQuit.ClickedCh

		srv.Shutdown(context.Background())
		systray.Quit()
	}()
}

func main() {
	if cfg.Headless {
		log.Fatal(srv.ListenAndServe())
	}

	systray.Run(onReady, func() {})
}
