package config

import (
	"errors"
	"fmt"

	"github.com/rodaine/table"
	"github.com/spf13/viper"
)

func printMap(m map[string]interface{}) {
	tbl := table.New("Key", "Value")
	printTable := true

	for key, value := range m {
		if str, ok := value.(string); ok {
			tbl.AddRow(key, str)
			continue
		}

		if subMap, ok := value.(map[string]interface{}); ok {
			printTable = false
			fmt.Println(key)
			printMap(subMap)
			fmt.Println()
		}
	}

	if printTable {
		tbl.Print()
	}
}

func Print() error {
	v := viper.New()
	v.AddConfigPath(getConfigDir())
	v.SetConfigType("yaml")
	v.SetConfigName("config")

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	printMap(v.AllSettings())

	return nil
}

func PrintField(field string) error {
	v := viper.New()
	v.AddConfigPath(getConfigDir())
	v.SetConfigType("yaml")
	v.SetConfigName("config")

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	value := v.Get(field)

	if _, ok := value.(map[string]interface{}); ok {
		printMap(value.(map[string]interface{}))
	}

	if _, ok := value.(string); ok {
		fmt.Println(value)
	}

	return nil
}

func SetField(field, value string) error {
	v := viper.New()
	v.AddConfigPath(getConfigDir())
	v.SetConfigType("yaml")
	v.SetConfigName("config")

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	switch field {
	case "general", "server", "ui", "ui.error", "ui.test":
		return errors.New("can only set primitive values")
	default:
		v.Set(field, value)
	}

	return v.WriteConfig()
}
