package config

import (
	"fmt"
	"os"

	"github.com/shelepuginivan/hakutest/internal/pkg/display"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

// options for ConfigService.
type options struct {
	v *viper.Viper
}

// Option represents constructor option for ConfigService.
type Option func(*options)

// WithViperInstance sets the underlying viper.Viper instance of the ConfigService.
func WithViperInstance(v *viper.Viper) Option {
	return func(o *options) {
		o.v = v
	}
}

// ConfigService is a struct that provides methods for manipulating configuration.
type ConfigService struct {
	options
}

// NewService returns a ConfigService instance.
func NewService(opts ...Option) *ConfigService {
	var options options

	for _, opt := range opts {
		opt(&options)
	}

	if options.v == nil {
		options.v = getViper()
	}

	return &ConfigService{options: options}
}

// PrintConfig prints the entire configuration.
func (s ConfigService) PrintConfig() error {
	if err := s.v.ReadInConfig(); err != nil {
		return err
	}

	display.PrintMap(s.v.AllSettings())

	return nil
}

// PrintField prints the specified configuration field.
// If the provided field represents a subbranch, it displays a table representation of this subbranch.
func (s ConfigService) PrintField(field string) error {
	if err := s.v.ReadInConfig(); err != nil {
		return err
	}

	value := s.v.Get(field)

	if branch, ok := value.(map[string]interface{}); ok {
		display.PrintMap(branch)
	} else {
		fmt.Println(value)
	}

	return nil
}

// GetField returns the specified field value.
func (s ConfigService) GetField(field string) any {
	if err := s.v.ReadInConfig(); err != nil {
		return nil
	}

	return s.v.Get(field)
}

// SetField sets the specified field value.
func (s ConfigService) SetField(field string, value any) error {
	if err := s.v.ReadInConfig(); err != nil {
		return err
	}

	s.v.Set(field, value)
	return s.v.WriteConfig()
}

// WriteConfig writes config cfg in a currectly used config file.
func (s ConfigService) WriteConfig(cfg *Config) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	if err := s.v.ReadInConfig(); err != nil {
		return err
	}

	file, err := os.OpenFile(s.v.ConfigFileUsed(), os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}
