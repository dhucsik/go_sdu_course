package controllers

import (
	"module/app/models"
	"module/pkg/utils"
	"module/platform/database"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body models.SignUp true "account info"
// @Success 200 {object} models.SignUp
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /sign-up [post]
func UserSignUp(c *fiber.Ctx) error {
	signUp := &models.SignUp{}

	if err := c.BodyParser(signUp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := utils.ValidateSignUp(signUp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	role, err := utils.VerifyRole(signUp.UserRole)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db := database.GetDatabase()

	user := &models.User{}

	user.Username = signUp.Username
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.Email = signUp.Email
	user.PasswordHash = utils.GeneratePassword(signUp.Password)
	user.UserRole = role

	if err := db.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	user.PasswordHash = ""

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  user,
	})
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body models.SignIn true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /sign-in [post]
func UserSignIn(c *fiber.Ctx) error {
	signIn := &models.SignIn{}

	if err := c.BodyParser(signIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := utils.ValidateSignIn(signIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db := database.GetDatabase()

	foundedUser, err := db.GetUserByEmail(signIn.Email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	compareUserPassword := utils.ComparePasswords(foundedUser.PasswordHash, signIn.Password)
	if !compareUserPassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "wrong user email address or password",
		})
	}

	tokens, err := utils.GenerateNewTokens(strconv.Itoa(foundedUser.UserID), foundedUser.UserRole)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"tokens": fiber.Map{
			"access":  tokens.Access,
			"refresh": tokens.Refresh,
		},
	})
}
