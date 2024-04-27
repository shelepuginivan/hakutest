package desktop

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/hakutest/internal/app/desktop/components"
	"github.com/shelepuginivan/hakutest/internal/app/desktop/layouts"
	"github.com/shelepuginivan/hakutest/internal/app/server"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
)

// ServerForm is a GTK component based on Gtk.Box.
// It provides a form to start the Hakutest server.
type ServerForm struct {
	*gtk.Box

	srv        *http.Server
	i18n       *i18n.GtkServerI18n
	srvRunning bool
}

// NewServerForm returns a new instance of ServerForm.
func (b Builder) NewServerForm() *ServerForm {
	form := &ServerForm{
		Box: Must(layouts.NewForm()),

		i18n:       b.app.I18n.Gtk.Server,
		srvRunning: false,
	}

	heading := Must(components.NewHeadingH1(form.i18n.Title))
	statusLabel := Must(gtk.LabelNew(form.i18n.LabelIdle))
	submitButton := Must(gtk.ButtonNewWithLabel(form.i18n.ButtonStart))

	submitButton.Connect("clicked", func(button *gtk.Button) {
		if form.srvRunning {
			form.srv.Shutdown(context.Background())
			form.srvRunning = false
			statusLabel.SetText(form.i18n.LabelIdle)
			button.SetLabel(form.i18n.ButtonStart)
			return
		}

		form.srv = server.New(b.app)

		go func() {
			if err := form.srv.ListenAndServe(); err != http.ErrServerClosed {
				form.srvRunning = false
				statusLabel.SetText(fmt.Sprintf(form.i18n.LabelError, err.Error()))
				button.SetLabel(form.i18n.ButtonStart)
			}
		}()

		form.srvRunning = true
		statusLabel.SetText(fmt.Sprintf(
			form.i18n.LabelRunning,
			b.app.Config.Server.Port,
		))
		button.SetLabel(form.i18n.ButtonStop)
	})

	form.PackStart(heading, false, false, 0)
	form.PackStart(submitButton, false, false, 0)
	form.PackStart(statusLabel, false, false, 0)

	return form
}
