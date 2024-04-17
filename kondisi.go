package main

import "fmt"

/*
If
if else
if else if
nested if
*/

func main() {
	nilaiUjian := 78
	nilaiAbsensi := 100

	if nilaiAbsensi > 90 && nilaiUjian > 95 {
		fmt.Println("Anda Lulus Dengan Nilai S+")
	} else if nilaiAbsensi > 80 && nilaiUjian > 90 {
		fmt.Println("Anda Lulus Dengan Nilai S")
	} else if nilaiAbsensi > 70 && nilaiUjian > 80 {
		fmt.Println("Anda Lulus Dengan Nilai A")
	} else if nilaiAbsensi > 60 && nilaiUjian > 70 {
		fmt.Println("Anda Lulus Dengan Nilai B")
	} else if nilaiAbsensi > 50 && nilaiUjian > 55 {
		fmt.Println("Anda Lulus Dengan Nilai C")
	} else {
		fmt.Println("Anda Tidak Lulus")
	}
}
