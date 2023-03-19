package controllers

import (
	"module/platform/database"
	"strconv"

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

func GetCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	category, err := db.GetCategory(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":    true,
			"msg":      "category with the given ID is not found",
			"category": nil,
		})
	}

	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      nil,
		"category": category,
	})
}
