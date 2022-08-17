package main

import "fmt"

// Create new type of 'deck'
// which is a type of strings
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}