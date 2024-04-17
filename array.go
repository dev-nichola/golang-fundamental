package main

import "fmt"

func main() {

	// Cara Pertama
	var lang [5]string

	// Menambahkan Array
	lang[0] = "Ruby"
	lang[1] = "Python"
	lang[2] = "PHP"
	lang[3] = "Javascript"
	lang[4] = "Java"

	// Cara Kedua
	bahasa := [5]string{
		"Jawa", "Batak", "Bali", "Mbojo", "Sunda",
	}

	// Cara Keempat
	hewan := [...]string{"Sapi", "Kucing", "Anjing", "Tak Terbatas"}
	// Cara Ketiga
	var kendaraan [5]string = [5]string{"Motor", "Mobil", "Becak", "Sepeda", "Bajaj"}
	fmt.Println(kendaraan)
	fmt.Println(bahasa)
	fmt.Println(lang)
	fmt.Println(hewan)
	fmt.Println("Jumlah Array Pada Variable Hewan Adalah : ", len(hewan))
	// Menentukan Panjang Array
	fmt.Println(len(lang))

	// Mengambil Array Satu Persatu
	fmt.Println(lang[1])

	// Mengambil Array Dengan Range
	for index, languagues := range lang {
		fmt.Println(index+1, languagues)
	}
}
