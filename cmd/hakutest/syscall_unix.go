//go:build !windows

package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
)

// sigusr registers callback for SIGUSR1.
func sigusr(_ *config.Config) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGUSR1)

	go func() {
		for {
			<-sig
			err := cfg.UpdateFromFile()
			if err != nil {
				log.Error().Err(err).Msg("Failed to update configuration from file")
			}
		}
	}()
}
