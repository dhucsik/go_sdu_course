package utils

import (
	"errors"
	"module/app/models"
)

func ValidateSignUp(model *models.SignUp) error {
	if len(model.Username) > 50 || len(model.Username) == 0 {
		return errors.New("invalid username")
	}

	if len(model.Email) > 255 || len(model.Email) == 0 {
		return errors.New("invalid email")
	}

	if len(model.Password) > 255 || len(model.Password) == 0 {
		return errors.New("invalid password")
	}

	return nil
}

func ValidateSignIn(model *models.SignIn) error {
	if len(model.Email) > 255 || len(model.Email) == 0 {
		return errors.New("invalid email")
	}

	if len(model.Password) > 255 || len(model.Password) == 0 {
		return errors.New("invalid password")
	}

	return nil
}

func ValidateProduct(model *models.Product) error {
	if len(model.ProductTitle) > 255 || len(model.ProductTitle) == 0 {
		return errors.New("invalid title")
	}

	if model.CategoryID < 0 {
		return errors.New("invalid category id")
	}

	if model.Price < 0 {
		return errors.New("invalid price")
	}

	if len(model.Description) > 255 || len(model.Description) == 0 {
		return errors.New("invalid description")
	}

	return nil
}

func ValidateReview(model *models.Review) error {
	if model.Rating < 0 || model.Rating > 5 {
		return errors.New("rating must be between 0 and 5")
	}

	return nil
}
