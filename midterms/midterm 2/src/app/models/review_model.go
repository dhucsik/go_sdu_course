package models

type Review struct {
	ReviewID int `db:"review_id" json:"review_id"`
	User     `json:"user"`
	Product  `json:"product"`
	Rating   int    `db:"rating" json:"rating"`
	Comment  string `db:"comment" json:"comment"`
}
