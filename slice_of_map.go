package main

import "fmt"

func main() {
	student := []map[string]string{
		{"name": "Nichola", "class": "Satu"},
		{"name": "Joko", "class": "Dua"},
		{"name": "Budi", "class": "Tiga"},
	}

	fmt.Println(student[0])

	for _, siswa := range student {
		fmt.Println(siswa)
	}
}
