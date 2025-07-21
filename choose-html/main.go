package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {

	type Chapter struct {
		Title     string   `json:"title"`
		Reference string   `json:"reference"`
		Story     []string `json:"story"`
		Options   []struct {
			Text    string `json:"text"`
			Chapter string `json:"arc"`
		} `json:"options"`
	}

	fileName := flag.String("file", "./gopher.json", "The JSON file for the story")
	flag.Parse()

	storyJSON, JSONErr := os.Open(*fileName)
	if JSONErr != nil {
		log.Fatal(JSONErr)
	}

	var story map[string]Chapter

	decodedStory := json.NewDecoder(storyJSON)
	if err := decodedStory.Decode(&story); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/intro", http.StatusMovedPermanently)
	})

	// The path value is the key that is accessed from the JSON
	http.HandleFunc("/{chapter}", func(w http.ResponseWriter, r *http.Request) {
		chapter := story[r.PathValue("chapter")]
		var sample Chapter

		if sample.Title == "" {
			http.Redirect(w, r, "/intro", http.StatusMovedPermanently)
		} else {
			template.Must(template.ParseFiles("temp.html")).Execute(w, chapter)
		}

	})

	http.ListenAndServe(":8080", nil)
}
