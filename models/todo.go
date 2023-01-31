package models

import (
	"github.com/jmoiron/sqlx"
)

type Todo struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
	Body  string `db:"body"`
	// Created_at string `db:"created_at"`
}

func FindAll(db *sqlx.DB) []Todo {
	todos := []Todo{}
	db.Select(&todos, `SELECT id, title, body FROM todo;`)
	return todos
}
func Add(db *sqlx.DB, title string, body string) {
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO todo (title, body) VALUES (?,?)", title, body)
	tx.Commit()
}
