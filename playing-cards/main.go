package main

func main() {
	// Let Go compiler determine type for the variable
	// Only used when defining a new variable
	// Creating a slice - expanding list with specific types
	cards := newDeck()

	cards.print()
}
