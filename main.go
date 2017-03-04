package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/kazuma-ickw/go-youtube/youtube"
)

var (
	query      = flag.String("query", "Google", "Search term")
	maxResults = flag.Int64("max-results", 25, "Max YouTube results")
)

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Printf("%#v", r)
		videos := youtube.Search(query, maxResults)
		for _, video := range videos {
			fmt.Printf("%#v", video)
		}
		fmt.Fprint(w, "hello world")
	})
	log.Fatal(http.ListenAndServe(":5000", nil))
}
