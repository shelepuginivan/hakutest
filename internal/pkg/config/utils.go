package config

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func printMap(m map[string]interface{}) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow, color.Bold).SprintfFunc()
	printKey := color.New(color.FgMagenta, color.Bold).PrintlnFunc()

	tbl := table.New("Key", "Value")
	tbl.WithHeaderFormatter(headerFmt)
	tbl.WithFirstColumnFormatter(columnFmt)

	printTable := true

	for key, value := range m {
		if subMap, ok := value.(map[string]interface{}); ok {
			printTable = false
			printKey(key)
			printMap(subMap)
			fmt.Println()
			continue
		}

		tbl.AddRow(key, value)
	}

	if printTable {
		tbl.Print()
	}
}

func Print() error {
	v := getViper()

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	printMap(v.AllSettings())

	return nil
}

func PrintField(field string) error {
	v := getViper()

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	value := v.Get(field)

	if _, ok := value.(map[string]interface{}); ok {
		printMap(value.(map[string]interface{}))
	} else {
		fmt.Println(value)
	}

	return nil
}

func SetField(field, value string) error {
	v := getViper()

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	switch field {
	case "general", "server", "stats", "stats.excel", "stats.image", "ui", "ui.editor", "ui.error", "ui.expired", "ui.test":
		return errors.New("can only set primitive values")
	default:
		v.Set(field, value)
	}

	return v.WriteConfig()
}
