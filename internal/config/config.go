package config

import (
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port           string `yaml:"port"`
	TestsDirectory string `yaml:"tests_directory"`
}

func getConfigPath() string {
	configDir, err := os.UserConfigDir()

	if err != nil {
		return "config.yaml"
	}

	return path.Join(configDir, "hakutest", "config.yaml")
}

func Init() Config {
	configPath := getConfigPath()
	config := Config{}
	port := "8080"
	testsDirectory := "user_test"
	cacheDir, err := os.UserCacheDir()

	if err == nil {
		testsDirectory = path.Join(cacheDir, "hakutest", "tests")
	}

	defaultConfig := Config{Port: port, TestsDirectory: testsDirectory}
	configFile, err := os.ReadFile(configPath)

	if err != nil {
		return defaultConfig
	}

	err = yaml.Unmarshal(configFile, &config)

	if err != nil {
		return defaultConfig
	}

	if config.TestsDirectory == "" {
		config.TestsDirectory = testsDirectory
	}

	if config.Port == "" {
		config.Port = port
	}

	return config
}
