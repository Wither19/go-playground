package main

import (
	"encoding/json"
	"log"
	"os"
)

// Reads JSON file named f and attempts to unmarshal its contents.
// Parsed JSON must be of struct T.
func JSONParse(f string) PathSetJSON {
	var paths PathSetJSON

	pathJSONFile, err := os.ReadFile(f)
	if err != nil {
		log.Fatalln("JSON file opening error:", err)
	}

	if err := json.Unmarshal(pathJSONFile, &paths); err != nil {
		log.Fatalln("JSON file unmarshal error:", err)	
	}

	return paths
}