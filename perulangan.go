package main

import (
	"fmt"
)

func main() {
	fmt.Println("Saya Sedang Belajar GOLANG")

	// for i := 0; i <= 100; i++ {
	// 	time.Sleep(time.Second)
	// 	if i == 10 {
	// 		fmt.Println("BLOK GOBLOK")
	// 	}
	// 	fmt.Println("Mencetak Perulangan Ke", i)
	// }

	// Range Itu Mirip Apa Di Foreach Kalau Di PHP
	// title := "Golang Adalah Bahasa Terhebat"
	// for _, letter := range title {
	// 	fmt.Println(string(letter))
	// }

	// Latihan
	title := "Golang Adalah Bahasa Terhebat"

	for index, letter := range title {
		if index%2 == 0 {
			fmt.Println("Index : ", index, "Letter : ", string(letter))
		}
	}
}
