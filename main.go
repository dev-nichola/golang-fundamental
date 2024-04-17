package main

import (
	"fmt"
	"golang-fundamental/calculation"
)

func main() {
	fmt.Println("==============")
	// cetak := TestCetak()
	// fmt.Println(cetak)

	result := calculation.Multiply(20, 10)

	fmt.Println("Hasil nya Ialah : ", result)
}
