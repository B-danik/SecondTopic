package service

type IBook interface {
	Get()
	Create()
}

type Manager struct {
	Book IBook
}

func NewManager() (*Manager, error) {
	return &Manager{}, nil
}
