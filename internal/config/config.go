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

func Init() Config {
	configPath := "config.yaml"
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

	return config
}
