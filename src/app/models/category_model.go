package models

type Category struct {
	ID    int64  `db:"id" json:"id" validate:"required"`
	Title string `db:"title" json:"title" validate:"required,lte=100"`
}
