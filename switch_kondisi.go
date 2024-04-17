package main

import "fmt"

func main() {
	number := 10

	switch number {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("Bukan Keduanya")
	}
}
