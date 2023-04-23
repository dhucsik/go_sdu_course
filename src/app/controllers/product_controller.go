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

/*
func GetProducts(c *fiber.Ctx) error {
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	products, err := db.GetProducts()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":    true,
			"msg":      err.Error(),
			"count":    0,
			"products": nil,
		})
	}

	//fmt.Println(products[0].SellerID.ID)
	//fmt.Println(products[0].CategoryID.ID)

	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      nil,
		"count":    len(products),
		"products": products,
	})
}
*/

// @Summary List Products
// @Tags product
// @Description list products
// @ID list-products
// @Accept json
// @Produce json
// @Success 200 {object} []models.Product
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /product [get]
func ListProducts(c *fiber.Ctx) error {
	queries := make(map[string]string, 5)
	queries["title"] = c.Query("name")
	queries["startPrice"] = c.Query("StartPrice")
	queries["endPrice"] = c.Query("EndPrice")
	queries["startRating"] = c.Query("StartRating")
	queries["endRating"] = c.Query("EndRating")

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	products, err := db.ListProducts(queries)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":    true,
			"msg":      err.Error(),
			"count":    0,
			"products": nil,
		})
	}

	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      nil,
		"count":    len(products),
		"products": products,
	})
}

// @Summary Get Product
// @Tags product
// @Description get product by ID
// @ID get-product
// @Accept json
// @Produce json
// @Param product_id path string true "product id"
// @Success 200 {object} models.Product
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /product/{product_id} [get]
func GetProduct(c *fiber.Ctx) error {
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

	product, err := db.GetProduct(id)
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
		"product": product,
	})
}

// @Summary Create Product
// @Security ApiKeyAuth
// @Tags product
// @Description create product
// @ID create-product
// @Accept  json
// @Produce  json
// @Param input body models.Product true "product info"
// @Success 200 {object} models.Product
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /product [post]
func CreateProduct(c *fiber.Ctx) error {
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

	userRole := claims.UserRole

	if userRole != repository.AdminRoleName && userRole == repository.SellerRoleName {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied",
		})
	}

	product := &models.Product{}

	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
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

	product.PasswordHash = ""

	if err := db.CreateProduct(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"product": product,
	})
}

// @Summary Update Product
// @Security ApiKeyAuth
// @Tags product
// @Description update product
// @ID update-product
// @Accept json
// @Produce json
// @Param product_id path string true "product id"
// @Param input body models.Product true "product info"
// @Success 200 {object} models.Product
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /product/{product_id} [put]
func UpdateProduct(c *fiber.Ctx) error {
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

	product := &models.Product{}

	if err := c.BodyParser(product); err != nil {
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

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	foundedProduct, err := db.GetProduct(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "product with the given ID is not found",
		})
	}

	userID := claims.UserID
	userRole := claims.UserRole

	if userRole != repository.AdminRoleName && userID != foundedProduct.UserID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied",
		})
	}

	if err := db.UpdateProduct(id, product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"product": product,
	})
}

// @Summary Delete Product
// @Security ApiKeyAuth
// @Tags product
// @Description delete product by ID
// @ID delete-product
// @Accept json
// @Produce json
// @Param product_id path string true "product id"
// @Success 200
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /product/{product_id} [delete]
func DeleteProduct(c *fiber.Ctx) error {
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

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	product, err := db.GetProduct(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "product with the given ID is not found",
		})
	}

	userRole := claims.UserRole
	userId := claims.UserID

	if userRole != repository.AdminRoleName && userId != product.UserID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied",
		})
	}

	if err := db.DeleteProduct(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
