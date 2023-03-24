package main

import (
	"os"

	"github.com/brandontor/booknookbackendgo/configs"
	"github.com/brandontor/booknookbackendgo/routes"
	"github.com/gofiber/fiber/v2"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()

	//routes
	routes.RouterSetup(app)

	app.Listen(getPort())

}
