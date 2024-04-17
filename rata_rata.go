package main

import "fmt"

func main() {
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	var total int = 0
	for _, score := range scores {
		total += score
	}

	ratarata := total / int(len(scores))

	fmt.Println(ratarata)

	var goodScores []int

	for _, score := range scores {
		if score >= 90 {
			goodScores = append(goodScores, score)
		}
	}

	fmt.Println(goodScores)
}
