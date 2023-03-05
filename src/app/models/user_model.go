package models

import "time"

type User struct {
	UserID       int       `db:"user_id" json:"user_id"`
	Username     string    `db:"username" json:"username" validate:"required,lte=50"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	Email        string    `db:"email" json:"email" validate:"required,email,lte=255"`
	PasswordHash string    `db:"password_hash" json:"password_hash" validate:"required,lte=255"`
	UserRole     string    `db:"user_role" json:"user_role" validate:"required,lte=25"`
}
