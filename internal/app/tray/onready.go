package tray

import (
	"github.com/getlantern/systray"
	"github.com/rs/zerolog/log"
)

type MenuEntry struct {
	Label    string
	Tooltip  string
	Callback func()
}

func OnReady(callback func(), entries ...MenuEntry) func() {
	return func() {
		systray.SetIcon(Icon)
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
