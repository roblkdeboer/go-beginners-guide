package main

import "fmt"

func main() {

	// Declaring the variable to put things in later
	// var colors map[string]string
	// colors := make(map[string]string)

	// Keys are type string
	// Values arae type string
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
	}

	colors["white"] = "#ffffff"

	// delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}