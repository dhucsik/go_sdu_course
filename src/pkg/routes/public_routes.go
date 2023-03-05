package routes

import (
	"module/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	route := app.Group("/api/v1")

	route.Get("/product", controllers.GetProducts)
	route.Get("/product/:id", controllers.GetProduct)

	route.Get("/category", controllers.GetCategories)
	route.Get("/category/:id", controllers.GetCategory)

	route.Post("/user/sign/up", controllers.UserSignUp)
	route.Post("/user/sign/in", controllers.UserSignIn)
}
