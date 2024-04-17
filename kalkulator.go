package main

import (
	"errors"
	"fmt"
)

func main() {
	tambah, err := kalkulator(100, 5, "/")

	if err != nil {
		panic(err)
	}

	fmt.Println(tambah)
}

func kalkulator(nilai1, nilai2 int, operator string) (int, error) {

	var result int
	var err error = errors.New("Unkown Operator")

	if operator == "+" {
		result = nilai1 + nilai2
	} else if operator == "-" {
		result = nilai1 - nilai2
	} else if operator == "*" {
		result = nilai1 * nilai2
	} else if operator == "/" {
		result = nilai1 / nilai2
	} else {
		return 0, err
	}

	return result, nil
}
