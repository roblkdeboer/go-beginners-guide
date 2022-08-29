package main

import "fmt"

type bot interface {
	// If anything has a function called getGreeting that returns a string, they are of type bot
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// printGreeting can now be used by bot types
func printGreeting(b bot) {
	fmt.Print(b.getGreeting())
}

// omit eb in function receiver when not used in function
func (englishBot) getGreeting() string {
	// custom logic for english greeting
	return "Hi there!\n"
}

func (spanishBot) getGreeting() string {
	// custom logic for spanish greeting
	return "Hola!"
}