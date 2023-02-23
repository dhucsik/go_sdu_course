package config

import (
	"database/sql"
	"module/pkg/db"
)

type application struct {
	User db.User
	DB   *sql.DB
}

var App application
