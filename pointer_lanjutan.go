package main

import "fmt"

func main() {
	var numberA int = 5
	var numberB *int = &numberA // & adalah proses refferensing

	fmt.Println(numberA)
	fmt.Println(*numberB)
	fmt.Println(numberB)

	numberA = 100
	fmt.Println(*numberB)
	fmt.Println(numberB)
	fmt.Println(numberA)
}
