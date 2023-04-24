package models

type SignUp struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserRole string `json:"user_role"`
}

type SignIn struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}
