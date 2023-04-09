package todo

type User struct {
	ID       int
	Name     string `json:name`
	LastName string `json:lastname`
	Email    string `json:email`
	Password string `json:password`
}
