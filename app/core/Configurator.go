package core

import (
	"encoding/json"
	"log"
	"os"
)

type Configurator struct {
}

func (c *Configurator) Read(path string, o *map[string]string) {

	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	err1 := decoder.Decode(o)

	if err1 != nil {
		log.Fatal(err1)
	}
}
