package controllers

import (
	"context"
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/rimo10/youtube-api-server/src/config"
	"github.com/rimo10/youtube-api-server/src/credentials"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
	"strconv"
)

var developerKey = credentials.API_KEY

func handleError(err error, str string) {
	if err != nil {
		log.Fatal(str)
	}
}

func Search(c *fiber.Ctx) error {
	flag.Parse()
	ctx := context.Background()

	query := c.Query("query", "")
	counts, err := strconv.ParseInt(c.Query("count", ""), 10, 64)

	if err != nil {
		return c.JSON("Incorrect query")
	}

	service, err := youtube.NewService(ctx, option.WithAPIKey(developerKey))
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Make the API call to YouTube.
	call := service.Search.List([]string{"id,snippet"}).
		Q(query).
		MaxResults(counts)
	response, err := call.Do()
	handleError(err, "Unable to fetch data from api")

	queryResponse := make([]*config.Searchapi, 0)

	for _, item := range response.Items {
		searchapi := &config.Searchapi{
			Query:       query,
			Etag:        item.Etag,
			VideoId:     item.Id.VideoId,
			Title:       item.Snippet.Title,
			Description: item.Snippet.Description,
			ChannelId:   item.Snippet.ChannelId,
			ChannelName: item.Snippet.ChannelTitle,
			PublishedAt: item.Snippet.PublishedAt,
		}
		queryResponse = append(queryResponse, searchapi)
	}

	for _, item := range queryResponse {
		if err := config.Database.Create(item).Error; err != nil {
			handleError(err, "Unable to add to database")
			return err
		}
	}

	return c.JSON(queryResponse)
}
