package main

import "fmt"

func main() {
	luas, _ := kalkulasi(10, 2)
	fmt.Println("Luasnya Adalah : ", luas)
}

func kalkulasi(panjang, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)

	return // Otomatis Ter Return
}
