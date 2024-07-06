// Package config provides configuration for the app.
package config

import (
	"os"
	"sync"

	"github.com/shelepuginivan/hakutest/internal/pkg/paths"
	"github.com/shelepuginivan/hakutest/pkg/security"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

// SecurityFields represents configuration for security policies.
type SecurityFields struct {
	Teacher string `yaml:"teacher"`
	Student string `yaml:"student"`
}

// Fields represents configuration fields.
type Fields struct {
	// General options.
	Debug    bool   `yaml:"debug"`    // Run in debug mode.
	Headless bool   `yaml:"headless"` // Run in headless mode (without systray icon).
	Port     int    `yaml:"port"`     // Port on which server is started.
	Lang     string `yaml:"lang"`

	// Results.
	OverwriteResults bool   `yaml:"overwrite_results"` // Whether to overwrite results on resend.
	ResultsDirectory string `yaml:"results_directory"`
	ShowResults      bool   `yaml:"show_results"` // Whether to show results on submission.

	// Security.
	Security SecurityFields `yaml:"security"`

	// Tests.
	TestsDirectory string `yaml:"tests_directory"`
}

// Config is a configuration layer for the application.
type Config struct {
	Fields

	callbacks []func(*Config)
	mu        sync.Mutex
}

// OnUpdate registers a callback allowing to run it when configuration is
// updated.
// This method is safe to use by multiple goroutines.
func (c *Config) OnUpdate(cb func(*Config)) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.callbacks = append(c.callbacks, cb)
}

// Update updates configuration fields and calls each registered callback.
// Provided Fields struct should contain all keys explicitly, otherwise
// unrepresented configuration fields are set to their zero value.
// This method is safe to use by multiple goroutines.
func (c *Config) Update(fields Fields) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Fields = fields

	for _, cb := range c.callbacks {
		cb(c)
	}
}

// New reads configuration file and returns the configuration.
// If field is unset, it fallbacks to the default value.
func New() *Config {
	cfg := Default()

	data, err := os.ReadFile(paths.Config)
	if err != nil {
		return Default()
	}

	if err = yaml.Unmarshal(data, &cfg.Fields); err != nil {
		return Default()
	}

	return cfg
}

// Default returns default configuration.
func Default() *Config {
	return &Config{
		Fields: Fields{
			Debug:            false,
			Headless:         false,
			Lang:             language.English.String(),
			Port:             8080,
			OverwriteResults: false,
			ResultsDirectory: paths.Results,
			Security: SecurityFields{
				Teacher: security.PolicyHostOnly,
				Student: security.PolicyNoVerification,
			},
			ShowResults:    true,
			TestsDirectory: paths.Tests,
		},
	}
}
