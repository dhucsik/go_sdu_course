package models

type Category struct {
	CategoryID    int    `db:"category_id" json:"category_id" validate:"required"`
	CategoryTitle string `db:"category_title" json:"category_title" validate:"required,lte=100"`
}
