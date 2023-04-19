package sql

import (
	"fmt"

	"github.com/B-danik/SecondTopic/todo"
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

func (b *Book) Create(name string, price int) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, price) values ($1, $2) RETURNING id", bookTable)

	row := b.db.QueryRow(query, name, price)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (b *Book) Get(ID int) (todo.Book, error) {

	var book todo.Book
	query := fmt.Sprintf("Select * from %s where id = $1", bookTable)

	err := b.db.Get(&book, query, ID)

	fmt.Print(book)
	return book, err
}

func (b *Book) GetList() ([]todo.Book, error) {

	var book []todo.Book
	query := fmt.Sprintf("Select * from %s", bookTable)

	err := b.db.Select(&book, query)

	return book, err
}

func (b *Book) Delete(ID int) error {

	query := fmt.Sprintf("Delete from %s where id = $1", bookTable)

	book, err := b.db.Exec(query, ID)

	fmt.Print(book)
	return err
}
