package main

import (
	"log"

	"module/pkg/middleware"
	"module/pkg/routes"

	"github.com/gofiber/fiber/v2"

	_ "module/docs"

	_ "github.com/joho/godotenv/autoload"
)

// @title 	E-commerce application
// @version	1.0
// @description E-commerce with jwt auth
// @termsOfService http://swagger.io/terms/

// @contact.name Dias Galikhanov
// @contact.email diasgalikhanov@gmail.com

// @host localhost:8586
// @BasePath /api/v1

// User + Auth CRUD + Swagger
// Product CRUD + Swagger

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app := fiber.New()

	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)

	if err := app.Listen(":8586"); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
