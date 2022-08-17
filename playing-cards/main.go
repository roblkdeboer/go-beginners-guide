package main

func main() {
	// Let Go compiler determine type for the variable
	// Only used when defining a new variable
	// Creating a slice - expanding list with specific types
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

// Expect the function to return string
func newCard() string {
	return "Five of Diamonds"
}