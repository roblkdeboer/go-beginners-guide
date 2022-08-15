package main

import "fmt"

func main() {

	// Explicitly type the variable
	// var card string = "Ace of Spades"

	// Let Go compiler determine type for the variable
	// Only used when defining a new variable
	card := "Ace of Spades"
	card = "Five of Diamonds"

	fmt.Println(card)
}