package main

import (
	"jv/url-shortener/maphandler"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", maphandler.MapHandler("paths.yml", "index.html"))
}
