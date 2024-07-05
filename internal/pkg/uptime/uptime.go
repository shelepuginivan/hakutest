// Package uptime provides method to get application uptime.
package uptime

import "time"

// Time of application start.
var startTime time.Time

func init() {
	startTime = time.Now()
}

// Uptime returns application uptime.
func Uptime() time.Duration {
	return time.Since(startTime)
}
