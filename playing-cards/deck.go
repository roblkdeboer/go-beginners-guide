package main

import "fmt"

// Create new type of 'deck'
// which is a type of strings
type deck []string

// Must return a value of type deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _,suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit )
		}
	}

	return cards
}

// d deck = receiver to receive a type deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}