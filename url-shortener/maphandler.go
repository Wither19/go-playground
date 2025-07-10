package main

import (
	"net/http"
	"strings"
	"text/template"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.

func stringRemove(str string, c string) string {
	return strings.ReplaceAll(str, c, "")
}

func MapHandler(pathsToUrls map[string]string) http.HandlerFunc {
	s := http.NewServeMux()

	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		shortenedPath := pathsToUrls[r.URL.Path]

		if (shortenedPath != "") {
			http.Redirect(w, r, shortenedPath, http.StatusFound)
		} else {
		template.Must(template.ParseFiles("index.html")).Funcs(template.FuncMap{
			"strRm": "stringRemove",
		}).Execute(w, pathsToUrls)
		}
})

	return s.ServeHTTP
}