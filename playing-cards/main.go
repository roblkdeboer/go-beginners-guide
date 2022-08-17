package main

func main() {
	// Let Go compiler determine type for the variable
	// Only used when defining a new variable
	// Creating a slice - expanding list with specific types
	cards := newDeck()

	// Gets assigned according to the order the decks are returned
	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()
}
