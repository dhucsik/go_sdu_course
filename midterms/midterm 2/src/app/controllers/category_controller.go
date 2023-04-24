package controllers

import (
	"module/platform/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// @Summary List Categories
// @Tags category
// @Description list categories
// @ID list-categories
// @Accept json
// @Produce json
// @Success 200 {object} []models.Category
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /category [get]
func GetCategories(c *fiber.Ctx) error {
	db := database.GetDatabase()

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

// @Summary Get Category
// @Tags category
// @Description get category by ID
// @ID get-category
// @Accept json
// @Produce json
// @Param category_id path string true "category id"
// @Success 200 {object} models.Category
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /category/{category_id} [get]
func GetCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db := database.GetDatabase()

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
