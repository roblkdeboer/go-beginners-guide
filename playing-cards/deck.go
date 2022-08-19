package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

// d deck = receiver to only work on a type deck, like a method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Input = type deck and handsize
// Return 2 decks
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Convert deck to strings
func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)

	if err != nil {
		// Option 1 - log the error and return a call to newDeck() - will always return a deck
		// Option 2 - log the error and quit the program
		log.Fatal("Error: ", err)
	}

	s := strings.Split(string(bs), ",")

	return deck(s)
}

func (d deck) shuffle() {
	// Need to generate a random number
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		// Swap elements in the slice
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}