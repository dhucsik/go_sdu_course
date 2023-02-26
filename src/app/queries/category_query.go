package queries

import (
	"module/app/models"

	"github.com/jmoiron/sqlx"
)

type CategoryQueries struct {
	*sqlx.DB
}

func (q *CategoryQueries) GetCategories() ([]models.Category, error) {
	categories := []models.Category{}

	query := `SELECT id, title FROM categories`

	err := q.Select(&categories, query)
	if err != nil {
		return categories, err
	}

	return categories, nil
}
