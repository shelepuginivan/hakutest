package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"fyne.io/systray"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/shelepuginivan/fsutil"
	"github.com/shelepuginivan/hakutest/internal/app/server"
	"github.com/shelepuginivan/hakutest/internal/app/tray"
	"github.com/shelepuginivan/hakutest/internal/pkg/browser"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/shelepuginivan/hakutest/internal/pkg/paths"
	"github.com/shelepuginivan/hakutest/pkg/result"
	"github.com/shelepuginivan/hakutest/pkg/security"
	"github.com/shelepuginivan/hakutest/pkg/test"
	"github.com/shelepuginivan/hakutest/pkg/version"
)

var cfg *config.Config
var displayVersion bool

func init() {
	cfg = config.New()

	flag.BoolVar(&displayVersion, "version", false, "Display version and exit")

	// General configuration.
	flag.BoolVar(&cfg.General.Debug, "general.debug", cfg.General.Debug, "Run in debug mode")
	flag.BoolVar(&cfg.General.DisableTray, "general.disable_tray", cfg.General.DisableTray, "Run without icon in system tray")
	flag.BoolVar(&cfg.General.OpenAtStartup, "general.open_at_startup", cfg.General.OpenAtStartup, "Open Hakutest web interface upon startup")
	flag.StringVar(&cfg.General.Lang, "general.lang", cfg.General.Lang, "Language of the interface")
	flag.IntVar(&cfg.General.Port, "general.port", cfg.General.Port, "Port on which server is started")

	// Result configuration.
	flag.StringVar(&cfg.Result.Path, "result.path", cfg.Result.Path, "Directory where results are stored")

	// Security configuration.
	flag.StringVar(&cfg.Security.Student, "security.student", cfg.Security.Student, "Security policy applied to student interface")
	flag.StringVar(&cfg.Security.Teacher, "security.teacher", cfg.Security.Teacher, "Security policy applied to teacher interface")
	flag.StringVar(&cfg.Security.Dialect, "security.dialect", cfg.Security.Dialect, "Dialect of the DB containing the user data")
	flag.StringVar(&cfg.Security.DSN, "security.dsn", cfg.Security.DSN, "DSN of the DB containing the user data")

	// Test configuration.
	flag.StringVar(&cfg.Test.Path, "test.path", cfg.Test.Path, "Directory where the test files are stored")
	flag.StringVar(&cfg.Test.DefaultTaskType, "test.default_task_type", cfg.Test.DefaultTaskType, "Default type of the new task added in the editor")
}

func onConfigUpdate(c *config.Config) {
	i18n.Init(c.General.Lang)
	result.Init(c.Result)
	test.Init(c.Test)
	initLogger(c.General.Debug)
}

func main() {
	flag.Parse()

	if displayVersion {
		fmt.Println(version.Version)
		os.Exit(0)
	}

	i18n.Init(cfg.General.Lang)
	result.Init(cfg.Result)
	security.Init(cfg.Security)
	test.Init(cfg.Test)
	initLogger(cfg.General.Debug)
	srv := server.New(cfg)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logFile, err := fsutil.CreateAll(paths.Logs)
	if err == nil {
		log.Logger = log.Output(logFile)
		defer logFile.Close()
	}

	cfg.OnUpdate(onConfigUpdate)

	// Update configuration when SIGUSR1 is sent.
	sigusr(cfg)

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
			Label:    i18n.Get("tray.open.label"),
			Tooltip:  i18n.Get("tray.open.tooltip"),
			Callback: browser.OpenDashboardFunc(cfg.General.Port),
		},
		tray.MenuEntry{
			Label:   i18n.Get("tray.quit.label"),
			Tooltip: i18n.Get("tray.quit.tooltip"),
			Callback: func() {
				srv.Shutdown(context.Background())
				systray.Quit()
				log.Info().Msg("Shutdown")
			},
		},
	)

	if cfg.General.OpenAtStartup {
		time.AfterFunc(1500*time.Millisecond, browser.OpenDashboardFunc(cfg.General.Port))
	}

	systray.Run(onReady, tray.OnExit)
}
