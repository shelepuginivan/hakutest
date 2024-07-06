package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/getlantern/systray"
	"github.com/shelepuginivan/hakutest/internal/app/server"
	"github.com/shelepuginivan/hakutest/internal/app/tray"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/shelepuginivan/hakutest/pkg/result"
	"github.com/shelepuginivan/hakutest/pkg/test"
)

var (
	cfg *config.Config
	srv *http.Server
)

func init() {
	cfg = config.New()

	flag.BoolVar(&cfg.Debug, "debug", cfg.Debug, "Run in debug mode")
	flag.BoolVar(&cfg.DisableTray, "disable-tray", cfg.DisableTray, "Run without icon in system tray")
	flag.StringVar(&cfg.Lang, "lang", cfg.Lang, "Language of the interface")
	flag.IntVar(&cfg.Port, "port", cfg.Port, "Port on which server is started")
	flag.StringVar(&cfg.TestsDirectory, "tests-directory", cfg.TestsDirectory, "Directory where the test files are stored")

	flag.Parse()

	i18n.Init(cfg.Lang)
	result.Init(cfg)
	test.Init(cfg.TestsDirectory)
	srv = server.New(cfg)
}

func main() {
	cfg.OnUpdate(func(c *config.Config) {
		i18n.Init(c.Lang)
		result.Init(c)
		test.Init(c.TestsDirectory)
	})

	if cfg.DisableTray {
		log.Fatal(srv.ListenAndServe())
	}

	onReady := tray.OnReady(
		func() {
			if err := srv.ListenAndServe(); err != nil {
				systray.Quit()
			}
		},
		tray.MenuEntry{
			Label:   "Quit",
			Tooltip: "Quit Hakutest",
			Callback: func() {
				srv.Shutdown(context.Background())
				systray.Quit()
			},
		},
	)

	systray.Run(onReady, tray.OnExit)
}
