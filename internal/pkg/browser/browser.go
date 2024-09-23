// Package browser provides utility functions that open teacher dashboard in the browser.
package browser

import (
	"fmt"

	"github.com/pkg/browser"
	"github.com/rs/zerolog/log"
)

// OpenDashboard opens teacher dashboard in the browser. If an error occurres,
// OpenDashboard logs it.
func OpenDashboard(port int) {
	err := browser.OpenURL(fmt.Sprintf("http://localhost:%d/teacher/dashboard", port))
	if err != nil {
		log.Error().Err(err).Msg("Failed to open dashboard")
	}
}

// OpenDashboardFunc is like [OpenDashboard], but returns a function. It is
// intended to be used in callbacks.
func OpenDashboardFunc(port int) func() {
	return func() {
		err := browser.OpenURL(fmt.Sprintf("http://localhost:%d/teacher/dashboard", port))
		if err != nil {
			log.Error().Err(err).Msg("Failed to open dashboard")
		}
	}
}
