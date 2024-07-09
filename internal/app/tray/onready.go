package tray

import (
	"fyne.io/systray"
	"github.com/rs/zerolog/log"
	"github.com/shelepuginivan/hakutest/internal/app/tray/icon"
)

type MenuEntry struct {
	Label    string
	Tooltip  string
	Callback func()
}

func OnReady(callback func(), entries ...MenuEntry) func() {
	return func() {
		systray.SetIcon(icon.Icon)
		systray.SetTitle(Title)
		systray.SetTooltip(Tooltip)

		for _, entry := range entries {
			menuItem := systray.AddMenuItem(entry.Label, entry.Tooltip)

			go func() {
				for {
					<-menuItem.ClickedCh
					entry.Callback()
				}
			}()
		}

		go callback()
		log.Info().Msg("System tray application is ready")
	}
}
