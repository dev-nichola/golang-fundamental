package main

import "fmt"

func main() {
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)

	fmt.Println(total)
}

func sum(nilai []int) int {

	total := 0
	for _, value := range nilai {
		total += value
	}

	return total
}
