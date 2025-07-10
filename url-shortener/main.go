package main

import (
	"net/http"
)

func main() {
	paths := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	http.ListenAndServe(":8080", MapHandler(paths))
}