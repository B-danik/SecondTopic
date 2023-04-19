package todo

type Book struct {
	Id    int
	Name  string
	Price int
}

type Rent struct {
	User_id int
	Book_id int
	Users   string
	Books   string
}
