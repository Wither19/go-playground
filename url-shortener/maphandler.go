package main

import (
	"net/http"
	"text/template"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string) http.HandlerFunc {
	
	s := http.NewServeMux()

	  s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			shortenedPath := pathsToUrls[r.URL.Path]

			if pathExists := shortenedPath != ""; pathExists {
				http.Redirect(w, r, shortenedPath, http.StatusFound)
			} else {
			template.Must(template.ParseFiles("index.html")).Execute(w, pathsToUrls)
			}
	})

	return s.ServeHTTP
}