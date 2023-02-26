package controllers

import (
	"module/platform/database"

	"github.com/gofiber/fiber/v2"
)

func GetCategories(c *fiber.Ctx) error {
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	categories, err := db.GetCategories()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":      true,
			"msg":        err.Error(),
			"count":      0,
			"categories": nil,
		})
	}

	return c.JSON(fiber.Map{
		"error":      false,
		"msg":        nil,
		"count":      len(categories),
		"categories": categories,
	})
}
