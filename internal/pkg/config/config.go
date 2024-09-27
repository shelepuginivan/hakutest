// Package config provides configuration for the app.
package config

import (
	"os"
	"sync"

	"github.com/shelepuginivan/fsutil"
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
	Debug bool `json:"debug" yaml:"debug"`

	// Run without icon in system tray.
	DisableTray bool `json:"disableTray" yaml:"disable_tray"`

	// Port on which server is started.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Language of the application interface.
	Lang string `json:"lang,omitempty" yaml:"lang,omitempty"`
}

// Fields represents configuration fields.
type Fields struct {
	// General configuration fields.
	General GeneralFields `json:"general,omitempty" yaml:"general,omitempty"`

	// Result package configuration.
	Result result.Config `json:"result,omitempty" yaml:"result,omitempty"`

	// Test package configuration.
	Test test.Config `json:"test,omitempty" yaml:"test,omitempty"`

	// Security configuration.
	Security security.Config `json:"security,omitempty" yaml:"security,omitempty"`
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

// UpdateFromFile updates configuration fields from the file and calls each
// registered callback.
//
// This method is safe to use by multiple goroutines.
func (c *Config) UpdateFromFile() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	f, err := read()
	if err != nil {
		return err
	}

	c.Fields = f
	for _, cb := range c.callbacks {
		cb(c)
	}

	return nil
}

// New reads configuration file and returns the configuration. If field is
// unset, it fallbacks to the default value as defined in [Default].
func New() *Config {
	f, err := read()
	if err != nil {
		return &Config{
			Fields: Default(),
		}
	}

	return &Config{
		Fields: f,
	}
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
func Default() Fields {
	return Fields{
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

// read reads configuration from the file.
func read() (Fields, error) {
	f := Default()

	data, err := os.ReadFile(paths.Config)
	if err != nil {
		return f, err
	}

	if err = yaml.Unmarshal(data, &f); err != nil {
		return f, err
	}

	return f, nil
}
