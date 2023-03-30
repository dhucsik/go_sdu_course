package queries

import (
	"module/app/models"

	"github.com/jmoiron/sqlx"
)

type ReviewQueries struct {
	*sqlx.DB
}

func (q *ReviewQueries) ListReviews(productID int) ([]models.Review, error) {
	reviews := []models.Review{}

	query := `SELECT r.review_id, 
				u.user_id, u.username, u.email,
				p.product_id, p.product_title,
				r.rating, r.comment
				FROM reviews r
				JOIN users u ON r.user_id = u.user_id
				JOIN products p ON r.product_id = p.product_id
				WHERE p.product_id = $1`

	err := q.Select(&reviews, query, productID)
	if err != nil {
		return reviews, err
	}

	return reviews, nil
}

func (q *ReviewQueries) CreateReview(r *models.Review) error {
	query := `INSERT INTO reviews(user_id, product_id, rating, comment) 
				VALUES ($1, $2, $3, $4)`

	_, err := q.Exec(query, r.UserID, r.ProductID, r.Rating, r.Comment)
	if err != nil {
		return err
	}

	return nil
}
