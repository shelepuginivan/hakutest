package main

import (
	"context"
	"flag"

	"fyne.io/systray"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/shelepuginivan/hakutest/internal/app/server"
	"github.com/shelepuginivan/hakutest/internal/app/tray"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/fsutil"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/shelepuginivan/hakutest/internal/pkg/paths"
	"github.com/shelepuginivan/hakutest/pkg/result"
	"github.com/shelepuginivan/hakutest/pkg/security"
	"github.com/shelepuginivan/hakutest/pkg/test"
)

var cfg *config.Config

func init() {
	cfg = config.New()

	flag.BoolVar(&cfg.General.Debug, "debug", cfg.General.Debug, "Run in debug mode")
	flag.BoolVar(&cfg.General.DisableTray, "disable-tray", cfg.General.DisableTray, "Run without icon in system tray")
	flag.StringVar(&cfg.General.Lang, "lang", cfg.General.Lang, "Language of the interface")
	flag.IntVar(&cfg.General.Port, "port", cfg.General.Port, "Port on which server is started")
	flag.StringVar(&cfg.Test.Path, "tests-directory", cfg.Test.Path, "Directory where the test files are stored")
}

func onConfigUpdate(c *config.Config) {
	i18n.Init(c.General.Lang)
	result.Init(c.Result)
	security.Init(c.Security)
	test.Init(c.Test)

	if c.General.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

func main() {
	flag.Parse()
	srv := server.New(cfg)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logFile, err := fsutil.CreateAll(paths.Logs)
	if err == nil {
		log.Logger = log.Output(logFile)
		defer logFile.Close()
	}

	cfg.OnUpdate(onConfigUpdate)
	cfg.Update(func(_ config.Fields) config.Fields {
		return cfg.Fields
	})

	if cfg.General.DisableTray {
		log.Fatal().Err(srv.ListenAndServe())
	}

	onReady := tray.OnReady(
		func() {
			if err := srv.ListenAndServe(); err != nil {
				systray.Quit()
				log.Fatal().Err(err)
			}
		},
		tray.MenuEntry{
			Label:   "Quit",
			Tooltip: "Quit Hakutest",
			Callback: func() {
				srv.Shutdown(context.Background())
				systray.Quit()
				log.Info().Msg("Shutdown")
			},
		},
	)

	systray.Run(onReady, tray.OnExit)
}
