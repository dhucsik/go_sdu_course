package models

import "time"

type User struct {
	UserID       int       `db:"user_id" json:"user_id"`
	Username     string    `db:"username" json:"username"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	Email        string    `db:"email" json:"email"`
	PasswordHash string    `db:"password_hash" json:"password_hash"`
	UserRole     string    `db:"user_role" json:"user_role"`
}
