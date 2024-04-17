package main

import "fmt"

func main() {
	result := tambah(10, 20)

	fmt.Println("Hasilnya Adalah : ", result)
}

func tambah(n1, n2 int) int {
	return n1 + n2
}
