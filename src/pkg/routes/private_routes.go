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
}
