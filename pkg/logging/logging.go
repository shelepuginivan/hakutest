// Package logging provides logging methods and helpers.
package logging

import (
	"log"
	"time"
)

// Println writes message to the log file.
// It forces format of message to match:
//
// `<prefix> | [info1 - info2 - info3 - ...] | <RFC1123 timestamp>\n`
func Println(prefix, info string) {
	log.Printf("%s | %s | %s\n", prefix, info, time.Now().Format(time.RFC1123))
}
