package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port string
}

func Init() Config {
	configPath := "config.yaml"
	configFile, err := os.ReadFile(configPath)
	config := Config{}

	if err != nil {
		return Config{Port: "8080"}
	}

	err = yaml.Unmarshal(configFile, &config)

	if err != nil {
		return Config{Port: "8080"}
	}

	return config
}
