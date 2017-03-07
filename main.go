package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/kazuma-ickw/go-youtube/youtube"
)

var (
	query      = flag.String("query", "Google", "Search term")
	maxResults = flag.Int64("max-results", 25, "Max YouTube results")
)

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%#v\n", r.URL.Query())
		fmt.Printf("%#v\n", r.URL.RequestURI())
		params := r.URL.Query()
		q := params.Get("q")
		max, _ := strconv.ParseInt(params.Get("max"), 10, 64)
		if max == 0 || max > 25 {
			max = 25
		}

		if len(q) == 0 {
			fmt.Printf("query not found")
		}

		videos := youtube.Search(q, max)
		jsonData, _ := json.Marshal(videos)
		w.Header().Set("Content-Type", "text/json")
		w.Write(jsonData)
	})
	log.Fatal(http.ListenAndServe(":5000", nil))
}
