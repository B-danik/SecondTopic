package sql

import "github.com/jmoiron/sqlx"

type Transactions struct {
	db *sqlx.DB
}

const (
	TransactionsTable = "books"
)

func NewTransactions(db *sqlx.DB) *Transactions {
	return &Transactions{db: db}
}
