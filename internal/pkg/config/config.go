package config

import (
	"os"
	"path/filepath"

	"github.com/shelepuginivan/hakutest/internal/pkg/directories"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type GeneralConfig struct {
	TestsDirectory   string `yaml:"tests_directory" mapstructure:"tests_directory"`
	ResultsDirectory string `yaml:"results_directory" mapstructure:"results_directory"`
	ShowResults      bool   `yaml:"show_results" mapstructure:"show_results"`
}

type ServerConfig struct {
	Port int    `yaml:"port" mapstructure:"port"`
	Mode string `yaml:"mode" mapstructure:"mode"`
}

type Config struct {
	General GeneralConfig `yaml:"general" mapstructure:"general"`
	Server  ServerConfig  `yaml:"server" mapstructure:"server"`
}

func getViper() *viper.Viper {
	v := viper.New()

	v.AddConfigPath(directories.Executable())
	v.AddConfigPath(directories.Config())
	v.SetConfigType("yaml")
	v.SetConfigName("config")

	return v
}

func Default() Config {
	dataDir := directories.Data()
	testsDirectory := filepath.Join(dataDir, "tests")
	resultsDirectory := filepath.Join(dataDir, "results")

	defaultConfig := Config{
		General: GeneralConfig{
			TestsDirectory:   testsDirectory,
			ResultsDirectory: resultsDirectory,
			ShowResults:      true,
		},
		Server: ServerConfig{
			Port: 8080,
			Mode: "release",
		},
	}

	return defaultConfig
}

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

	return config
}
