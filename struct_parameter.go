package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	email     string
	isActive  bool
}

func main() {
	user := displayUser(User{
		FirstName: "Nichola",
		LastName:  "Saputra",
		email:     "nicholasaputra77@gmail.com",
	})

	fmt.Println(user)
}

func displayUser(user User) string {
	result := fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.email)
	return result
}
