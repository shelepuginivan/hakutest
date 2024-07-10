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
	"github.com/shelepuginivan/hakutest/pkg/test"
)

var cfg *config.Config

func init() {
	cfg = config.New()

	flag.BoolVar(&cfg.Debug, "debug", cfg.Debug, "Run in debug mode")
	flag.BoolVar(&cfg.DisableTray, "disable-tray", cfg.DisableTray, "Run without icon in system tray")
	flag.StringVar(&cfg.Lang, "lang", cfg.Lang, "Language of the interface")
	flag.IntVar(&cfg.Port, "port", cfg.Port, "Port on which server is started")
	flag.StringVar(&cfg.TestsDirectory, "tests-directory", cfg.TestsDirectory, "Directory where the test files are stored")
}

func onConfigUpdate(c *config.Config) {
	i18n.Init(c.Lang)
	result.Init(c)
	test.Init(c)

	if c.Debug {
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
	cfg.Update(cfg.Fields)

	if cfg.DisableTray {
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
