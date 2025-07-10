package main

import (
	"net/http"
	"text/template"

	"github.com/Masterminds/sprig"
)

// pathsToUrls is a string map containing aliases to shorten URL
// addresses. The key is the shortened link and the corresponding
// value is the full URL it is aliased to. If the given request is
// not using a URL alias, the server will instead load the fallback
// page containing a list of the aliases.
func MapHandler() http.HandlerFunc {
	pathsToUrls := yamlParse("paths.yml")
	s := http.NewServeMux()

	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		shortenedPath := pathsToUrls[r.URL.Path]

		if (shortenedPath != "") {
			http.Redirect(w, r, shortenedPath, http.StatusFound)
		} else {
			temp := template.New("index.html")
			temp = temp.Funcs(sprig.FuncMap())

			template.Must(temp.ParseFiles("index.html")).Execute(w, pathsToUrls)
		}
})

	return s.ServeHTTP
}