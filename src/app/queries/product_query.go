package queries

import (
	"fmt"
	"module/app/models"

	"github.com/jmoiron/sqlx"
)

type ProductQueries struct {
	*sqlx.DB
}

/*
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

		return products, nil
	}
*/
func (q *ProductQueries) ListProducts(queries map[string]string) ([]models.Product, error) {
	products := []models.Product{}

	if queries["startPrice"] == "" {
		queries["startPrice"] = "0"
	}
	if queries["endPrice"] == "" {
		queries["endPrice"] = "10000000"
	}

	query := `SELECT p.product_id, p.product_title, 
				u.user_id, u.username, u.email, 
				c.category_id, c.category_title, 
				p.price, p.description, p.avg_rating
				FROM products p 
				JOIN users u ON p.seller_id = u.user_id
				JOIN categories c ON p.category_id = c.category_id
				WHERE LOWER(p.product_title) LIKE $1
				AND p.price >= $2 AND p.price <= $3
				AND p.avg_rating >= $4
				AND p.avg_rating <= $5`

	err := q.Select(&products, query, fmt.Sprintf("%%%s%%", queries["title"]), queries["startPrice"], queries["endPrice"], queries["startRating"], queries["endRating"])
	if err != nil {
		return products, err
	}

	return products, nil
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

func (q *ProductQueries) UpdateProduct(id int, p *models.Product) error {
	query := `UPDATE products SET product_title = $2, category_id = $3, price = $4, description = $5
				WHERE product_id = $1`

	_, err := q.Exec(query, id, p.ProductTitle, p.CategoryID, p.Price, p.Description)
	if err != nil {
		return err
	}

	return nil
}

func (q *ProductQueries) DeleteProduct(id int) error {
	query := `DELETE FROM products WHERE product_id = $1`

	_, err := q.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
