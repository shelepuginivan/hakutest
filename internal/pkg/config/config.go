// Package config provides configuration for the app.
package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config is a global application configuration layer.
type Config struct {
	Debug    bool `yaml:"debug"`    // Run in debug mode.
	Headless bool `yaml:"headless"` // Run in headless mode (without systray icon).
	Port     int  `yaml:"port"`     // Port on which server is started.
}

// New reads configuration file and returns the configuration.
// If field is unset, it fallbacks to the default value.
func New() *Config {
	cfg := Default()

	data, err := os.ReadFile(configFile())
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
		Debug:    false,
		Headless: false,
		Port:     8080,
	}
}
