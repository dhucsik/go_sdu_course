package routes

import (
	"module/app/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func PublicRoutes(app *fiber.App) {
	route := app.Group("/api/v1")

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:          "http://example.com/doc.json",
		DeepLinking:  false,
		DocExpansion: "none",
	}))

	route.Get("/product", controllers.ListProducts)
	route.Get("/product/:id", controllers.GetProduct)

	route.Get("/product/:id/review", controllers.ListReviews)
	route.Get("/review/:id", controllers.GetReview)

	route.Get("/category", controllers.GetCategories)
	route.Get("/category/:id", controllers.GetCategory)

	route.Post("/user/sign/up", controllers.UserSignUp)
	route.Post("/user/sign/in", controllers.UserSignIn)
}
