package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rimo10/youtube-api-server/db"
	"github.com/rimo10/youtube-api-server/src/routes"
)



func main() {
	app := fiber.New()
	routes.SetSearchRoutes(app)
	app.Listen(":3000")
	defer db.GetDB().Close()
}
