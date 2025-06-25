package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"golang.org/x/net/html"
)


func main() {
	fileFlag := flag.String("file", "example.html", "The HTML file to read")
	flag.Parse()

	htmlOpen, htmlOpenErr := os.ReadFile(*fileFlag)
	if (htmlOpenErr != nil) {
		log.Fatal(htmlOpenErr)
	}
	htmlDoc := string(htmlOpen)

	parsedDoc, parseErr := html.Parse(strings.NewReader(htmlDoc))

	if (parseErr != nil) {
		log.Fatal(parseErr)
	}

	for t := range parsedDoc.Descendants() {
		if (slices.Contains(t.Attr, html.Attribute{ Key: "href" })) {
			fmt.Println("yes")
		}
	}
}