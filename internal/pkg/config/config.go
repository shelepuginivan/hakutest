// Package config provides configuration for the app.
package config

import (
	"os"

	"github.com/shelepuginivan/hakutest/internal/pkg/paths"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

// Config is a global application configuration layer.
type Config struct {
	// General options.
	Debug    bool   `yaml:"debug"`    // Run in debug mode.
	Headless bool   `yaml:"headless"` // Run in headless mode (without systray icon).
	Port     int    `yaml:"port"`     // Port on which server is started.
	Lang     string `yaml:"lang"`

	// Results.
	OverwriteResults bool   `yaml:"overwrite_results"` // Whether to overwrite results on resend.
	ResultsDirectory string `yaml:"results_directory"`
	ShowResults      bool   `yaml:"show_results"` // Whether to show results on submission.

	// Tests.
	TestsDirectory string `yaml:"tests_directory"`
}

// New reads configuration file and returns the configuration.
// If field is unset, it fallbacks to the default value.
func New() *Config {
	cfg := Default()

	data, err := os.ReadFile(paths.Config)
	if err != nil {
		return Default()
	}

	if err = yaml.Unmarshal(data, cfg); err != nil {
		return Default()
	}

	return cfg
}

// Default returns default configuration.
func Default() *Config {
	return &Config{
		Debug:            false,
		Headless:         false,
		Lang:             language.English.String(),
		Port:             8080,
		OverwriteResults: false,
		ResultsDirectory: paths.Results,
		ShowResults:      true,
		TestsDirectory:   paths.Tests,
	}
}
