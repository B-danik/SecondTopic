package sql

import (
	"fmt"

	"github.com/B-danik/SecondTopic/todo"
	"github.com/jmoiron/sqlx"
)

type Rent struct {
	db *sqlx.DB
}

const (
	rentTable = "rent_list"
)

func NewRent(db *sqlx.DB) *Rent {
	return &Rent{db: db}
}

func (r *Rent) CraeteRent(user_id int, book_id int) error {

	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, book_id) values ($1, $2) RETURNING id", rentTable)

	row := r.db.QueryRow(query, user_id, book_id)
	if err := row.Scan(&id); err != nil {
		return err
	}

	return nil
}

func (r *Rent) GetRent(id int) ([]todo.Rent, error) {

	var rent []todo.Rent

	query := fmt.Sprintf("SELECT p.name AS users, b.name AS books FROM users p"+
		" LEFT JOIN rent_list pb ON pb.user_id = p.id"+
		" LEFT JOIN books b ON b.id = pb.book_id where p.id = %d;", id)

	// query := fmt.Sprintf("SELECT p.name AS $1, b.name AS $2 FROM $1 p" +
	// 	" LEFT JOIN $3 pb ON pb.user_id = p.id" +
	// 	" LEFT JOIN $2 b ON b.id = pb.book_id;")

	err := r.db.Select(&rent, query)

	fmt.Print(rent)
	return rent, err
}
