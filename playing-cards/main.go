package main

import "fmt"

func main() {

	// Explicitly type the variable
	// var card string = "Ace of Spades"

	// Let Go compiler determine type for the variable
	// Only used when defining a new variable
	card := newCard()

	fmt.Println(card)
}

// Expect the function to return string
func newCard() string {
	return "Five of Diamonds"
}