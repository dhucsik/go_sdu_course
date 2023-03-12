package queries

import (
	"module/app/models"

	"github.com/jmoiron/sqlx"
)

type UserQueries struct {
	*sqlx.DB
}

func (q *UserQueries) CreateUser(u *models.User) error {
	query := `INSERT INTO users(username, created_at, email, password_hash, user_role) VALUES ($1, $2, $3, $4, $5)`

	_, err := q.Exec(
		query,
		u.Username, u.CreatedAt, u.Email, u.PasswordHash, u.UserRole,
	)
	if err != nil {
		return err
	}

	return nil
}

func (q *UserQueries) GetUserByEmail(email string) (models.User, error) {
	user := models.User{}

	query := `SELECT * FROM users WHERE email = $1`

	err := q.Get(&user, query, email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (q *UserQueries) GetUserByID(id int) (models.User, error) {
	user := models.User{}

	query := `SELECT * FROM users WHERE user_id = $1`

	err := q.Get(&user, query, id)
	if err != nil {
		return user, err
	}

	return user, nil
}
