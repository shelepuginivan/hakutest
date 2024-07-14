// Package config provides configuration for the app.
package config

import (
	"os"
	"sync"

	"github.com/shelepuginivan/hakutest/internal/pkg/fsutil"
	"github.com/shelepuginivan/hakutest/internal/pkg/paths"
	isecurity "github.com/shelepuginivan/hakutest/internal/pkg/security"
	"github.com/shelepuginivan/hakutest/pkg/result"
	"github.com/shelepuginivan/hakutest/pkg/security"
	"github.com/shelepuginivan/hakutest/pkg/test"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

// GeneralFields represents general configuration fields defined in `general`
// section in the configuration file.
type GeneralFields struct {
	// Run in debug mode.
	Debug bool `yaml:"debug"`

	// Run without icon in system tray.
	DisableTray bool `yaml:"disable_tray"`

	// Port on which server is started.
	Port int `yaml:"port"`

	// Language of the application interface.
	Lang string `yaml:"lang"`
}

// Fields represents configuration fields.
type Fields struct {
	// General configuration fields.
	General GeneralFields `yaml:"general"`

	// Result package configuration.
	Result result.Config `yaml:"result"`

	// Test package configuration.
	Test test.Config `yaml:"test"`

	// Security configuration.
	Security security.Config `yaml:"security"`
}

// Config is a configuration layer for the application.
type Config struct {
	Fields

	callbacks []func(*Config)
	mu        sync.Mutex
}

// OnUpdate registers a callback allowing to run it when configuration is
// updated.
//
// This method is safe to use by multiple goroutines.
func (c *Config) OnUpdate(cb func(*Config)) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.callbacks = append(c.callbacks, cb)
}

// Update updates configuration fields and calls each registered callback.
// Provided Fields struct should contain all keys explicitly, otherwise
// unrepresented configuration fields are set to their zero value.
//
// This method is safe to use by multiple goroutines.
func (c *Config) Update(updateFunc func(f Fields) Fields) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Fields = updateFunc(c.Fields)

	for _, cb := range c.callbacks {
		cb(c)
	}

	return write(c)
}

// New reads configuration file and returns the configuration. If field is
// unset, it fallbacks to the default value as defined in [Default].
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

// Default returns the default configuration:
//
//	general:
//	  debug: false
//	  disable_tray: false
//	  port: 8080
//	  lang: en
//	result:
//	  overwrite: false
//	  path: $XDG_DATA_HOME/hakutest/results
//	  show: true
//	test:
//	  path: $XDG_DATA_HOME/hakutest/tests
//	security:
//	  dsn: $XDG_CACHE_HOME/hakutest/users.db
//	  dialect: sqlite
//	  teacher: hostonly
//	  student: no_verification
func Default() *Config {
	return &Config{
		Fields: Fields{
			General: GeneralFields{
				Debug:       false,
				DisableTray: false,
				Lang:        language.English.String(),
				Port:        8080,
			},
			Result: result.Config{
				Overwrite: false,
				Path:      paths.Results,
				Show:      true,
			},
			Security: security.Config{
				DSN:     paths.UserDB,
				Dialect: isecurity.DialectSQLite,
				Teacher: security.PolicyHostOnly,
				Student: security.PolicyNoVerification,
			},
			Test: test.Config{
				Path: paths.Tests,
			},
		},
	}
}

// write writes configuration to the file.
func write(cfg *Config) error {
	data, err := yaml.Marshal(cfg.Fields)
	if err != nil {
		return err
	}

	return fsutil.WriteAll(paths.Config, data)
}
