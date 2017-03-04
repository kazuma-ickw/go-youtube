package youtube

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/BurntSushi/toml"

	"google.golang.org/api/googleapi/transport"
	youtube "google.golang.org/api/youtube/v3"
)

// Config is a root-level configration object
type Config struct {
	API APIConfig
}

// APIConfig is a configration object for youtube api
type APIConfig struct {
	Key string `toml:"key"`
}

// Video is a video object
type Video struct {
	Title        string
	ID           string
	ThumbnailURL string
	PublishedAt  time.Time
}

var (
	query      = flag.String("query", "Google", "Search term")
	maxResults = flag.Int64("max-results", 25, "Max YouTube results")
)

// Search Video
func Search() []Video {
	flag.Parse()

	var config Config
	_, err := toml.DecodeFile("config.toml", &config)

	client := &http.Client{
		Transport: &transport.APIKey{Key: config.API.Key},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}
	call := service.Search.List("id,snippet").
		Q(*query).
		MaxResults(*maxResults)

	response, err := call.Do()
	if err != nil {
		// The channels.list method call returned an error.
		log.Fatalf("Error making API call to list channels: %v", err.Error())
	}

	var videos []Video

	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			t, _ := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
			videos = append(videos, Video{
				Title:        item.Snippet.Title,
				ID:           item.Id.VideoId,
				ThumbnailURL: item.Snippet.Thumbnails.Default.Url,
				PublishedAt:  t,
			})
		}
	}

	return videos
}
