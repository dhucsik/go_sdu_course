package routes

import (
	"module/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	route := app.Group("/api")

	route.Get("/categories", controllers.GetCategories)
}
