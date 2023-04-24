package controllers

import (
	"module/app/models"
	"module/pkg/repository"
	"module/pkg/utils"
	"module/platform/database"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// @Summary List Reviews
// @Tags review
// @Description list reviews
// @ID list-reviews
// @Accept json
// @Produce json
// @Param product_id path int true "product id"
// @Success 200 {object} []models.Review
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /product/{product_id}/review [get]
func ListReviews(c *fiber.Ctx) error {
	productID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db := database.GetDatabase()

	reviews, err := db.ListReviews(productID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"msg":     err.Error(),
			"count":   0,
			"reviews": nil,
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"count":   len(reviews),
		"reviews": reviews,
	})
}

// @Summary Create Review
// @Security ApiKeyAuth
// @Tags review
// @Description create review
// @ID create-review
// @Accept  json
// @Produce  json
// @Param input body models.Review true "review info"
// @Success 200 {object} models.Review
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /product/:id/review [post]
func CreateReview(c *fiber.Ctx) error {
	now := time.Now().Unix()

	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	expires := claims.Expires

	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	productID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	review := &models.Review{}

	if err := c.BodyParser(review); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	review.ProductID = productID
	review.UserID = claims.UserID

	db := database.GetDatabase()

	if err := db.CreateReview(review); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := db.UpdateProductAvgRating(review.ProductID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"review": review,
	})

	/*
		product, err := db.GetProduct(productID)
		if err != nil {
			return c.Status(c.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg": err.Error(),
			})
		}

		user, err := db.GetUserByID(claims.UserID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
	*/

}

// @Summary Get Review
// @Tags review
// @Description get review by ID
// @ID get-review
// @Accept json
// @Produce json
// @Param review_id path string true "review id"
// @Success 200 {object} models.Review
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /review/{review_id} [get]
func GetReview(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db := database.GetDatabase()

	review, err := db.GetReview(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"msg":     "product with the given ID is not found",
			"product": nil,
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"product": review,
	})
}

// @Summary Update Review
// @Security ApiKeyAuth
// @Tags review
// @Description update review
// @ID update-review
// @Accept json
// @Produce json
// @Param review_id path string true "review id"
// @Param input body models.Review true "review info"
// @Success 200 {object} models.Review
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /review/{review_id} [put]
func UpdateReview(c *fiber.Ctx) error {
	now := time.Now().Unix()

	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	expires := claims.Expires

	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	review := models.Review{}

	if err := c.BodyParser(review); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db := database.GetDatabase()

	foundedReview, err := db.GetReview(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "review with the given ID not found",
		})
	}

	userID := claims.UserID
	userRole := claims.UserRole

	if userRole != repository.AdminRoleName && userID != foundedReview.UserID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied",
		})
	}

	if err := db.UpdateReview(id, review); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := db.UpdateProductAvgRating(review.ProductID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"review": review,
	})
}

// @Summary Delete Review
// @Security ApiKeyAuth
// @Tags review
// @Description delete review by ID
// @ID delete-review
// @Accept json
// @Produce json
// @Param review_id path string true "review id"
// @Success 200
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /review/{review_id} [delete]
func DeleteReview(c *fiber.Ctx) error {
	now := time.Now().Unix()

	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	expires := claims.Expires

	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db := database.GetDatabase()

	review, err := db.GetReview(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "review with the given ID is not found",
		})
	}

	userRole := claims.UserRole
	userId := claims.UserID

	if userRole != repository.AdminRoleName && userId != review.UserID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied",
		})
	}

	if err := db.DeleteReview(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
