// Package display provides methods for printing different data types.
package display

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

// PrintMap prints a table representation of a map data structure.
func PrintMap(m map[string]interface{}) {
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
			PrintMap(subMap)
			fmt.Println()
			continue
		}

		tbl.AddRow(key, value)
	}

	if printTable {
		tbl.Print()
	}
}
