package tray

import "github.com/shelepuginivan/hakutest/pkg/logging"

func OnExit() {
	logging.Println("TRAY", "System tray application stopped")
}
