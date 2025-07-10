package main

import (
	"log"
	"os"

	"github.com/stretchr/testify/assert/yaml"
)

// Reads YAML file named f and attempts to unmarshal its contents.
// Parsed YAML must be a string map.
func yamlParse(f string) map[string]string {
	var paths map[string]string

	pathYamlFile, err := os.ReadFile(f)
	if err != nil {
		log.Fatalln("YAML file opening error:", err)
	}

	if err := yaml.Unmarshal(pathYamlFile, &paths); err != nil {
		log.Fatalln("YAML file unmarshal error:", err)	
	}

	return paths
}