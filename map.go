package main

import "fmt"

func main() {
	var makanan map[string]string = map[string]string{
		"Ruby": "Is Beautiful",
		"GO":   "Super Fast",
		"PHP":  "Lambo",
	}

	for index, lang := range makanan {
		fmt.Println(index, lang)
	}

	fmt.Println("==================")
	delete(makanan, "GO")
	for index, lang := range makanan {
		fmt.Println(index, lang)
	}
}
