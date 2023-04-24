package routes

import (
	"module/app/controllers"
	"module/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/product", middleware.JWTProtected(), controllers.CreateProduct)
	route.Delete("/product/:id", middleware.JWTProtected(), controllers.DeleteProduct)
	route.Put("/product/:id", middleware.JWTProtected(), controllers.UpdateProduct)

	route.Post("/product/:id/review", middleware.JWTProtected(), controllers.CreateReview)
	route.Put("/review/:id", middleware.JWTProtected(), controllers.UpdateReview)
	route.Delete("/review/:id", middleware.JWTProtected(), controllers.DeleteReview)
}
