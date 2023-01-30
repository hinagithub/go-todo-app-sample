package models

type Todo struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
	Body  string `db:"body"`
}
