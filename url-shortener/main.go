package main

import (
	"log"
	"net/http"
	"os"

	"github.com/stretchr/testify/assert/yaml"
	"gopkg.in/yaml.v3"
)

func main() {
	
	var paths map[string]string

	pathYamlFile, err := os.ReadFile("paths.yml")
	if err != nil {
		log.Fatalln("YAML file opening error:", err)
	}

	if err := yaml.Unmarshal(pathYamlFile, paths); err != nil {
		log.Fatalln("YAML file unmarshal error:", err)	
	}

	http.ListenAndServe(":8080", MapHandler(paths))
}