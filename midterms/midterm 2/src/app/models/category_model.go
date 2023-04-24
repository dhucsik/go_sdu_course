package models

type Category struct {
	CategoryID    int    `db:"category_id" json:"category_id"`
	CategoryTitle string `db:"category_title" json:"category_title"`
}
