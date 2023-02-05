package models

import (
	"go-api/util"

	"github.com/jmoiron/sqlx"
)

type Todo struct {
	ID        int    `db:"id"`
	Completed bool   `db:"completed"`
	Body      string `db:"body"`
}

func FindAll(db *sqlx.DB) []Todo {
	todos := []Todo{}
	db.Select(&todos, `SELECT id, completed, body FROM todo WHERE deleted_at IS NULL`)
	return todos
}
func Add(db *sqlx.DB, completed bool, body string) {
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO todo (completed, body) VALUES (?,?)", completed, body)
	tx.Commit()
}
func Edit(db *sqlx.DB, id int, completed bool, body string) error {
	todo := Todo{}
	err := db.Get(&todo, "SELECT id, completed, body FROM todo WHERE id=? AND deleted_at IS NULL", id)
	if err != nil {
		return err
	}
	tx := db.MustBegin()
	tx.MustExec("UPDATE todo SET completed=?, body=? WHERE id=?", completed, body, id)
	tx.Commit()
	return nil
}
func Delete(db *sqlx.DB, id int) error {
	todo := Todo{}
	err := db.Get(&todo, "SELECT id, completed, body FROM todo WHERE id=? AND deleted_at IS NULL", id)
	if err != nil {
		return err
	}
	tx := db.MustBegin()
	now := util.Now()
	tx.MustExec("UPDATE todo SET deleted_at=? WHERE id=?", now, id)
	tx.Commit()
	return nil
}
