package main

import (
	"github.com/brandontor/booknookbackendgo/configs"
	"github.com/brandontor/booknookbackendgo/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()

	//routes
	routes.RouterSetup(app)

	app.Listen(":3000")

}
