package routes

import (
	"github.com/brandontor/booknookbackendgo/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouterSetup(app *fiber.App) {

	booknooks := app.Group("/booknooks")
	user := app.Group("/user")

	//All routes related to booknooks come here
	booknooks.Post("/", controllers.CreateBookNook) //add this
	booknooks.Get("/", controllers.GetBookNookList)
	booknooks.Get("/:bookNookId", controllers.GetBookNook)

	user.Get("/", func(c *fiber.Ctx) error {
		return nil
	})

}
