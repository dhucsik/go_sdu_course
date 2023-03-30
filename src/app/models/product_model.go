package models

type Product struct {
	ProductID    int    `db:"product_id" json:"product_id"`
	ProductTitle string `db:"product_title" json:"product_title"`
	User         `json:"user"`
	Category     `json:"category"`
	Price        float64 `db:"price" json:"price"`
	Description  string  `db:"description" json:"description"`
	AvgRating    float64 `db:"avg_rating" json:"avg_rating"`
}
