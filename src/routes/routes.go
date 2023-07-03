package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rimo10/youtube-api-server/src/controllers"
)

var SetSearchRoutes = func(app *fiber.App) {
	app.Get("/api/search_videos", controllers.Search) //searching videos from the api
	app.Get("/api/get_videos", controllers.Get)       //getting videos from the database
}
