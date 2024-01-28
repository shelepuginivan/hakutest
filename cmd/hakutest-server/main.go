package main

import (
	"os"

	"github.com/getlantern/systray"
	"github.com/shelepuginivan/hakutest/internal/app/server"
	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
)

func main() {
	systray.Run(onReady, func() {})
}

func getIcon(s string) []byte {
	b, err := os.ReadFile(s)
	if err != nil {
		panic(err)
	}
	return b
}

func onReady() {
	systray.SetTitle("Hakutest server")
	systray.SetTooltip("Hakutest")
	systray.SetIcon(getIcon("assets/hakutest.ico"))

	mStop := systray.AddMenuItem("Stop Hakutest", "Stop Hakutest server and quit")

	go server.NewRouter(test.NewService(), results.NewService()).Run()

	go func() {
		<-mStop.ClickedCh
		systray.Quit()
	}()
}
