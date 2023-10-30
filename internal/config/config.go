package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port           string `yaml:"port"`
	TestsDirectory string `yaml:"tests_directory"`
}

func Init() Config {
	configPath := "config.yaml"
	config := Config{}
	port := "8080"
	testDirectory, err := os.UserCacheDir()

	if err != nil {
		testDirectory = "user_tests"
	}

	defaultConfig := Config{Port: port, TestsDirectory: testDirectory}
	configFile, err := os.ReadFile(configPath)

	if err != nil {
		return defaultConfig
	}

	err = yaml.Unmarshal(configFile, &config)

	if err != nil {
		return defaultConfig
	}

	return config
}
