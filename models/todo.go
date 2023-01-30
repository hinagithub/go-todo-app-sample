package models

import (
	"github.com/jmoiron/sqlx"
)

func GetAll(db *sqlx.DB) []Todo {
	todos := []Todo{}
	db.Select(&todos, `SELECT id, title, body FROM todo;`)
	return todos
}
