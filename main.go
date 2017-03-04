package main

import (
	"fmt"

	"github.com/kazuma-ickw/go-youtube/youtube"
)

func main() {
	videos := youtube.Search()
	fmt.Println("return")

	for _, video := range videos {
		fmt.Printf("%#v", video)
	}
}
