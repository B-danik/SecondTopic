package service

type IUsers interface {
	Get()
	Create()
}

type logUser struct {
	Email    string `json:email`
	Password string `json:password`
}

type Service struct {
	User IUsers
}

func NewService() (*Service, error) {
	return &Service{}, nil
}
