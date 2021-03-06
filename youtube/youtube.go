package youtube

import (
	"fmt"
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
	Title        string    `json:"title"`
	ID           string    `json:"id"`
	URL          string    `json:"url"`
	ThumbnailURL string    `json:"thumbnailUrl"`
	PublishedAt  time.Time `json:"publishedAt"`
}

// Search Video
func Search(query string, maxResults int64) []Video {

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
		Q(query).
		MaxResults(maxResults + 1)

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
				URL:          fmt.Sprintf("https://www.youtube.com/watch?v=%v", item.Id.VideoId),
				ThumbnailURL: item.Snippet.Thumbnails.Default.Url,
				PublishedAt:  t,
			})
		}
	}

	return videos
}
