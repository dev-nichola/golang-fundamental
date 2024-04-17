package main

import (
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	email     string
	isActive  bool
}

func main() {
	user := User{
		FirstName: "Nichola",
		LastName:  "Saputra",
		email:     "nichola@gmail.com",
	}
	result := user.display()

	fmt.Println(result)
}

func (user User) display() string {
	return fmt.Sprintf("Nama %s %s, Email: %s", user.FirstName, user.LastName, user.email)
}
