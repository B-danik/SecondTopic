package book

import (
	repository "github.com/B-danik/SecondTopic/internal/database/postgre"
)

type BookService struct {
	repo repository.IBook
}

func NewBook(repo repository.IBook) *BookService {
	return &BookService{repo: repo}
}

func (b *BookService) CreateBook(name string) (int, error) {
	b.repo.CreateBook(name)
	return 0, nil

}

func (b *BookService) GetBook() {

}
