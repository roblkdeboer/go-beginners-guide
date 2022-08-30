package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
	}

	// Communicate through the channel with strings
	// Can send data across channels to communicate between channels
	c := make(chan string) 

	for _, link := range links {
		// Each go routine is a set of code executing inside a function
		// go keyword means the a separate go routine is launched to run the code in the function
		go checkLink(link, c)
	}
	// Receive the message from the channel
	// This is a blocking function so we need to loop infinitely to keep checking the sites
	// For link in the channel:
	for l := range c {

		// Function literal -> anonymous function to execute some code and invoke it with a (l) to pass the link into the function
		// Link gets copied in memory so the child routine can refer to that copied link
		go func(link string) {
		// Sleep should not be in main routine as the messages get throttled
		// Sleep should not be in checkLink routine as we want to check the sites immediately
		time.Sleep(5*time.Second)
		checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	// Get request is a blocking function
	// When a blocking function is hit, control flow allows the next go routine to take over
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		// Send message to channel
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// Send message to channel
	c <- link
}