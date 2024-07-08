package tray

import "github.com/rs/zerolog/log"

func OnExit() {
	log.Info().Msg("System tray application exited")
}
