package main

import "fmt"

func main() {
	luas, _ := calculate(10, 2)
	fmt.Println("Luasnya Adalah : ", luas)
}

func calculate(panjang, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	return luas, keliling
}
