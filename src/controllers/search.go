package controllers

import (
	"context"
	"flag"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rimo10/youtube-api-server/src/config"
	"github.com/rimo10/youtube-api-server/src/credentials"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var developerKey = credentials.API_KEY

func Search(c *fiber.Ctx) error {
	flag.Parse()
	ctx := context.Background()

	query := c.Query("query", "")
	counts, err := strconv.ParseInt(c.Query("count", ""), 10, 64)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "unable to fetch results for your given query. Make sure to type the correct query",
		})
	}

	service, err := youtube.NewService(ctx, option.WithAPIKey(developerKey))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "error in fetching data from the api",
		})
	}

	// Make the API call to YouTube.
	call := service.Search.List([]string{"id,snippet"}).
		Q(query).
		MaxResults(counts)
	response, err := call.Do()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "your query was accepted but api couldn't fetch the data",
		})
	}
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
			return c.Status(http.StatusBadRequest).JSON(map[string]string{
				"error": "unable to add item to the database",
			})
		}
	}

	return c.JSON(queryResponse)
}
