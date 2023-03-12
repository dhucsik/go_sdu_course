package main

import (
	"log"

	"module/pkg/middleware"
	"module/pkg/routes"

	"github.com/gofiber/fiber/v2"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New()

	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)

	if err := app.Listen(":8080"); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
