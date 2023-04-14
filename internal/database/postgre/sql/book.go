package sql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Book struct {
	db *sqlx.DB
}

const (
	bookTable = "books"
)

func NewBook(db *sqlx.DB) *Book {
	return &Book{db: db}
}

func (b *Book) CreateBook(name string) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name) values ($1) RETURNING id", bookTable)

	row := b.db.QueryRow(query, name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (b *Book) GetBook() {

}
