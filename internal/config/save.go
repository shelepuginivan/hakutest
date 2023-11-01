package config

import (
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

func (c Config) Save() error {
	configPath := getConfigPath()
	data, err := yaml.Marshal(c)

	if err != nil {
		log.Fatal(err)
	}

	err = os.MkdirAll(path.Dir(configPath), 0770)

	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	return os.WriteFile(configPath, data, 0666)
}
