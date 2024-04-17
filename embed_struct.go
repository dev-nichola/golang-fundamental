package main

import "fmt"

// Struct Attribute Struct lainya
type User struct {
	ID        int
	FirstName string
	LastName  string
	email     string
	isActive  bool
}

type Group struct {
	Name        string
	Admin       User
	users       []User
	IsAvailable bool
}

func main() {
	user1 := User{
		ID:        1,
		FirstName: "Nichola",
		LastName:  "Saputra",
		email:     "nichola@gmail.com",
		isActive:  true,
	}

	user2 := User{
		ID:        2,
		FirstName: "Joko",
		LastName:  "Morro",
		email:     "joko@gmail.com",
		isActive:  false,
	}

	users := []User{user1, user2}
	group := Group{"Gamer", user1, users, true}

	fmt.Println(group)
}
