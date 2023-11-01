package parser

import (
	"encoding/json"
	"log"
	"os"
)

func (t Test) Save(name string) error {
	testPath := GetTestPath(name)
	data, err := json.Marshal(t)

	if err != nil {
		log.Fatal(err)
	}

	return os.WriteFile(testPath, data, 0666)
}
