package main

import (
	"fmt"
)

func main() {
	number := 5
	fmt.Println("Alamat Memori : ", &number)
	fmt.Println("Nilai Awal :", number)

	change(&number, 100)

	fmt.Println("Nilai Akhir :", number)
	fmt.Println("Alamat Memori : ", &number)
}

func change(old *int, newValue int) {
	*old = newValue
	fmt.Println("Di Dalam Function :", *old, newValue)
	fmt.Println("Alamat Memori : ", &old)

}
