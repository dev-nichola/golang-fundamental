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

	// Cara Pengisian Pertama
	user1 := User{
		ID:        1,
		FirstName: "Nichola",
		LastName:  "Saputra",
		email:     "nichola@gmail.com",
		isActive:  true,
	}

	// Cara Pengisian Kedua
	user2 := User{}
	user2.ID = 2
	user2.FirstName = "Joko"
	user2.LastName = "Morro"
	user2.email = "jokomorro@gmail.com"
	user2.isActive = false

	// Cara Pengisian Ketiga
	user3 := User{
		3, "Budi", "Handoko", "budi@gmail.com", true,
	}

	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user3)

}
