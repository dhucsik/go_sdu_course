package models

type Product struct {
	ProductID    int    `db:"product_id" json:"product_id"`
	ProductTitle string `db:"product_title" json:"product_title" validate:"required,lte=100"`
	User         `json:"user"`
	Category     `json:"category"`
	Price        float64 `db:"price" json:"price" validate:"required"`
	Description  string  `db:"description" json:"description" validate:"required,lte=300"`
}
