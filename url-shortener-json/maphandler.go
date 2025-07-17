package main

import (
	"html/template"
	"net/http"

	"github.com/Masterminds/sprig"
)


type PathSetJSON map[string]struct{
	Path string `json:"path"`
	HTMLHidden bool `json:"htmlHidden"`
}

// pathsToUrls is a string map containing aliases to shorten URL
// addresses. The key is the shortened link and the corresponding
// value is the full URL it is aliased to. If the given request is
// not using a URL alias, the server will instead load the fallback
// page containing a list of the aliases.
func MapHandler(paths string, fallbackPage string) http.HandlerFunc {
	pathsToUrls := JSONParse(paths)
	s := http.NewServeMux()

	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		shortenedPath := pathsToUrls[r.URL.Path]

		if (shortenedPath.Path != "") {
			http.Redirect(w, r, shortenedPath.Path, http.StatusFound)
		} else {
			temp := template.New("index.html")
			temp = temp.Funcs(sprig.FuncMap())

			template.Must(temp.ParseFiles(fallbackPage)).Execute(w, pathsToUrls)
		}
})

	return s.ServeHTTP
}