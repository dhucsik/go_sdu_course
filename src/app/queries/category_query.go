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

	query := `SELECT * FROM categories`

	err := q.Select(&categories, query)
	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (q *CategoryQueries) GetCategory(id int) (models.Category, error) {
	category := models.Category{}

	query := `SELECT * FROM categories WHERE category_id = $1`

	err := q.Get(&category, query, id)
	if err != nil {
		return category, err
	}

	return category, nil
}
