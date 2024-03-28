package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rimo10/youtube-api-server/db"
	"github.com/rimo10/youtube-api-server/src/routes"
)

func main() {
	app := fiber.New()
	routes.SetSearchRoutes(app)
	app.Listen(":3000")
	defer db.CloseDB()
}
