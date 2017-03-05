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
		fmt.Printf("%#v", r.URL.Query())
		params := r.URL.Query()
		q := params.Get("q")
		fmt.Printf("%#v", q)

		if len(q) == 0 {
			fmt.Printf("query not found")
		}

		videos := youtube.Search(q, maxResults)
		for _, video := range videos {
			fmt.Printf("%#v", video)
		}
		fmt.Fprint(w, "hello world")
	})
	log.Fatal(http.ListenAndServe(":5000", nil))
}
