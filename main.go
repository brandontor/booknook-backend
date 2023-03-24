package main

import (
	"log"
	"os"

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

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))

}
