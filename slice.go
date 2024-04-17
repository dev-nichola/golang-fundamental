package main

import "fmt"

func main() {

	// Cara Pertama
	fish := []string{
		"Salmon",
		"Hiu",
		"Paus",
		"Tongkol",
	}

	for i, fishs := range fish {
		fmt.Println(i+1, fishs)
	}

	// Cara Kedua
	var gamingConsole []string

	// Menambahkan Slice
	gamingConsole = append(gamingConsole, "PS2")

	fmt.Println(gamingConsole)
}
