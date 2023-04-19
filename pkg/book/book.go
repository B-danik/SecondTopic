package book

import (
	repository "github.com/B-danik/SecondTopic/internal/database/postgre"
	"github.com/B-danik/SecondTopic/todo"
)

type BookService struct {
	repo repository.IBook
}

func NewBook(repo repository.IBook) *BookService {
	return &BookService{repo: repo}
}

func (b *BookService) Create(name string, price int) (int, error) {
	return b.repo.Create(name, price)

}

func (b *BookService) Get(ID int) (todo.Book, error) {
	return b.repo.Get(ID)
}

func (b *BookService) GetList() ([]todo.Book, error) {
	return b.repo.GetList()
}

func (b *BookService) Delete(ID int) error {
	return b.repo.Delete(ID)
}
