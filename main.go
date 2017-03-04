package main

import (
	"flag"
	"fmt"

	"github.com/kazuma-ickw/go-youtube/youtube"
)

var (
	query      = flag.String("query", "Google", "Search term")
	maxResults = flag.Int64("max-results", 25, "Max YouTube results")
)

func main() {
	flag.Parse()
	videos := youtube.Search(query, maxResults)
	fmt.Println("return")

	for _, video := range videos {
		fmt.Printf("%#v", video)
	}
}
