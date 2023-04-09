package authUser

// import (
// 	"log"
// 	"os"
// )

// type User struct {
// 	Name     string
// 	Lastname string
// 	Email    string
// 	Password string
// }

// // Временно использывал файл место базы данных
// func NewFile(name string, lastname string, email string, password string) *User {
// 	return &User{
// 		Name:     name,
// 		Lastname: lastname,
// 		Email:    email,
// 		Password: password,
// 	}
// }

// func (s *User) CreateFile() {
// 	log.Println("Create or update database...")
// 	file, err := os.Create(s.Email + ".txt")

// 	if err != nil {
// 		log.Println("Unable to create file:", err)
// 		os.Exit(1)
// 	}
// 	defer file.Close()
// 	file.WriteString(s.Name + "\n")
// 	file.WriteString(s.Lastname + "\n")
// 	file.WriteString(s.Email + "\n")
// 	file.WriteString(s.Password + "\n")
// 	log.Println("Add user: ", s)
// 	log.Println("Done.")
// }

// func ReadFile(mail string) (string, error) {
// 	log.Print("Read database...")
// 	file, err := os.Open(mail + ".txt")
// 	if err != nil {
// 		log.Println(err)
// 		return "", nil
// 	}
// 	defer file.Close()
// 	data := make([]byte, 64)
// 	n, err := file.Read(data)
// 	return string(data[:n]), nil
// }
