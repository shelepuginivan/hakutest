//go:build windows

package main

import "github.com/shelepuginivan/hakutest/internal/pkg/config"

// sigusr registers callback for SIGUSR1.
//
// Since Windows does not implement SIGUSR1, the function is empty.
func sigusr(_ *config.Config) {}
