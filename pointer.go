package main

import "fmt"

func main() {
	numberA := 5
	numberB := &numberA

	// Number A dan Number B mengacu ke titik tempat yang sama
	fmt.Println(*numberB)

	*numberB = 10
	fmt.Println(*numberB)
	fmt.Println(numberA) // Kita Langsung Mengubah
}
