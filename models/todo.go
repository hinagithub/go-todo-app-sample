package models

import (
	"github.com/jmoiron/sqlx"
)

type Todo struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
	Body  string `db:"body"`
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
func Edit(db *sqlx.DB, id int, title string, body string) error {
	todo := Todo{}
	err := db.Get(&todo, "SELECT id, title, body FROM todo WHERE id=?", id)
	if err != nil {
		return err
	}
	tx := db.MustBegin()
	tx.MustExec("UPDATE todo SET title=?, body=? where id=?", title, body, id)
	tx.Commit()
	return nil
}
