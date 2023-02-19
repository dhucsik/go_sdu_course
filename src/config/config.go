package config

import "module/pkg/db"

type application struct {
	User db.User
}

var App application
