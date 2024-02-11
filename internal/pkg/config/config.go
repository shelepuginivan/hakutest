// Package config provides global configuration for Hakutest.
package config

import (
	"os"
	"path/filepath"

	"github.com/shelepuginivan/hakutest/internal/pkg/directories"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

// GeneralConfig represents general configuration parameters.
type GeneralConfig struct {
	// Directory where tests are stored (environment variables such as $HOME are supported).
	TestsDirectory string `yaml:"tests_directory" mapstructure:"tests_directory"`

	// Directory where results are stored (environment variables such as $HOME are supported).
	ResultsDirectory string `yaml:"results_directory" mapstructure:"results_directory"`

	// Specifies whether the results will be displayed immediately after the response is sent.
	ShowResults bool `yaml:"show_results" mapstructure:"show_results"`

	// Specifies whether the results are allowed to be overwritten if the same student resubmits the solution again.
	OverwriteResults bool `yaml:"overwrite_results" mapstructure:"overwrite_results"`
}

// ServerConfig represents server configuration parameters.
type ServerConfig struct {
	// Port on which server is started.
	Port int `yaml:"port" mapstructure:"port"`

	// Mode in which server is started.
	Mode string `yaml:"mode" mapstructure:"mode"`
}

// Config represents Hakutest configuration.
type Config struct {
	General GeneralConfig `yaml:"general" mapstructure:"general"` // General configuration.
	Server  ServerConfig  `yaml:"server" mapstructure:"server"`   // Server configuration.
}

// getViper returns a configured instance of viper.Viper.
// It scans the OS-specific configuration directory and the Hakutest executable directory for the configuration file `config.yaml`.
func getViper() *viper.Viper {
	v := viper.New()

	v.AddConfigPath(directories.Executable())
	v.AddConfigPath(directories.Config())
	v.SetConfigType("yaml")
	v.SetConfigName("config")

	return v
}

// Default returns the default configuration.
func Default() Config {
	dataDir := directories.Data()
	testsDirectory := filepath.Join(dataDir, "tests")
	resultsDirectory := filepath.Join(dataDir, "results")

	defaultConfig := Config{
		General: GeneralConfig{
			TestsDirectory:   testsDirectory,
			ResultsDirectory: resultsDirectory,
			ShowResults:      true,
			OverwriteResults: false,
		},
		Server: ServerConfig{
			Port: 8080,
			Mode: "release",
		},
	}

	return defaultConfig
}

// createDefaultConfig creates a configuration file in the OS-specific configuration directory.
func createDefaultConfig() error {
	configDir := directories.Config()
	configPath := filepath.Join(configDir, "config.yaml")

	err := os.MkdirAll(configDir, os.ModeDir|os.ModePerm)

	if err != nil {
		return err
	}

	file, err := os.Create(configPath)

	if err != nil {
		return err
	}

	defer file.Close()

	data, err := yaml.Marshal(Default())

	if err != nil {
		data = []byte{}
	}

	_, err = file.Write(data)

	return err
}

// New returns configuration defined in the configuration file.
// Fields that are not specified in the configuration file are fallback to default values.
func New() Config {
	config := Default()

	v := getViper()
	v.SetDefault("general", config.General)
	v.SetDefault("server", config.Server)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(err)
		}

		if err := createDefaultConfig(); err != nil {
			panic(err)
		}
	}

	if err := v.Unmarshal(&config); err != nil {
		panic(err)
	}

	config.General.TestsDirectory = os.ExpandEnv(config.General.TestsDirectory)
	config.General.ResultsDirectory = os.ExpandEnv(config.General.ResultsDirectory)

	return config
}
