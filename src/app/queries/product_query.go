package queries

import (
	"module/app/models"

	"github.com/jmoiron/sqlx"
)

type ProductQueries struct {
	*sqlx.DB
}

func (q *ProductQueries) GetProducts() ([]models.Product, error) {
	products := []models.Product{}

	query := `SELECT p.product_id, p.product_title, 
				u.user_id, u.username, u.email, 
				c.category_id, c.category_title, 
				price, description 
				FROM products p 
				JOIN users u ON p.seller_id = u.user_id
				JOIN categories c ON p.category_id = c.category_id`

	err := q.Select(&products, query)
	if err != nil {
		return products, err
	}

	return products, err
}

func (q *ProductQueries) GetProduct(id int) (models.Product, error) {
	product := models.Product{}

	query := `SELECT p.product_id, p.product_title, 
				u.user_id, u.username, u.email, 
				c.category_id, c.category_title, 
				price, description 
				FROM products p 
				JOIN users u ON p.seller_id = u.user_id
				JOIN categories c ON p.category_id = c.category_id
				WHERE p.product_id = $1`

	err := q.Get(&product, query, id)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (q *ProductQueries) CreateProduct(p *models.Product) error {
	query := `INSERT INTO products(product_title, seller_id, category_id, price, description) 
				VALUES ($1, $2, $3, $4, $5)`

	_, err := q.Exec(query, p.ProductTitle, p.UserID, p.CategoryID, p.Price, p.Description)
	if err != nil {
		return err
	}

	return nil
}
