package rent

import (
	repository "github.com/B-danik/SecondTopic/internal/database/postgre"
	"github.com/B-danik/SecondTopic/todo"
)

type Rent struct {
	repo repository.IRent
}

func NewRent(repo repository.IRent) *Rent {
	return &Rent{repo: repo}
}

func (r *Rent) GetRent(id int) ([]todo.Rent, error) {
	return r.repo.GetRent(id)
}
func (r *Rent) CraeteRent(user_id int, book_id int) error {
	return r.repo.CraeteRent(user_id, book_id)
}
