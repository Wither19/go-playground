package urlshortener

import (
	"maps"
	"net/http"
	"slices"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	Keys of the map are the shortened URLs of their values.

	// URL map example from the repo:
	// paths := map[string]string{
	// "/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	// "/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	// }
	
	keys := slices.Collect(maps.Keys(pathsToUrls))

	s := http.NewServeMux()

	for url, i := range keys {
		s.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
			http.RedirectHandler()
		})
	}
	
	

	return nil
}