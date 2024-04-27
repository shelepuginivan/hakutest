package desktop

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/hakutest/internal/app/desktop/components"
	"github.com/shelepuginivan/hakutest/internal/app/desktop/layouts"
	"github.com/shelepuginivan/hakutest/internal/app/server"
)

// ServerForm is a GTK component based on Gtk.Box.
// It provides a form to start the Hakutest server.
type ServerForm struct {
	*gtk.Box

	srv        *http.Server
	srvRunning bool
}

// NewServerForm returns a new instance of ServerForm.
func (b Builder) NewServerForm() *ServerForm {
	form := &ServerForm{
		Box: Must(layouts.NewForm()),

		srvRunning: false,
	}

	heading := Must(components.NewHeadingH1("Hakutest Server"))
	statusLabel := Must(gtk.LabelNew("Server is not running"))
	submitButton := Must(gtk.ButtonNewWithLabel("Start server"))

	submitButton.Connect("clicked", func(button *gtk.Button) {
		if form.srvRunning {
			form.srv.Shutdown(context.Background())
			statusLabel.SetText("Server is not running")
			button.SetLabel("Start server")
			return
		}

		form.srv = server.New(b.app)
		form.srvRunning = true

		go func() {
			if err := form.srv.ListenAndServe(); err != http.ErrServerClosed {
				statusLabel.SetText("Server is not running")
				button.SetLabel("Start server")
			}
		}()

		statusLabel.SetText(fmt.Sprintf(
			"Server is running on port :%d",
			b.app.Config.Server.Port,
		))
		button.SetLabel("Stop server")
	})

	form.PackStart(heading, false, false, 0)
	form.PackStart(submitButton, false, false, 0)
	form.PackStart(statusLabel, false, false, 0)

	return form
}
