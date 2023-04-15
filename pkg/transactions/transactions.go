package transactions

import (
	repository "github.com/B-danik/SecondTopic/internal/database/postgre"
)

type TransactionsService struct {
	repo repository.ITransactions
}

func NewBook(repo repository.ITransactions) *TransactionsService {
	return &TransactionsService{repo: repo}
}
