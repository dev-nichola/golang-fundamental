package main

import "fmt"

type Gamer struct {
	Name  string
	Games []string
}

// Add Game
func (gamer *Gamer) AddGame(game string) {
	gamer.Games = append(gamer.Games, game)
}
func main() {
	gamer := Gamer{Name: "Nichola"}

	gamer.AddGame("Zelda")
	gamer.AddGame("Boku No Hero Academia")
	gamer.AddGame("GTA SA")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}
