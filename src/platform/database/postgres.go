package database

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"module/app/queries"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

var database *Queries

func GetDatabase() *Queries {
	return database
}

type Queries struct {
	*queries.UserQueries
	*queries.ProductQueries
	*queries.CategoryQueries
	*queries.ReviewQueries
}

func OpenDBConnection() error {
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	url := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL_MODE"),
	)

	db, err := sqlx.Connect("pgx", url)
	if err != nil {
		return fmt.Errorf("error, not connected to database, %w", err)
	}

	_, err = db.Exec(migration)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	if err := db.Ping(); err != nil {
		defer db.Close()
		return fmt.Errorf("error, not sent ping to database, %w", err)
	}

	database = &Queries{
		UserQueries:     &queries.UserQueries{DB: db},
		ProductQueries:  &queries.ProductQueries{DB: db},
		CategoryQueries: &queries.CategoryQueries{DB: db},
		ReviewQueries:   &queries.ReviewQueries{DB: db},
	}

	return nil
}
