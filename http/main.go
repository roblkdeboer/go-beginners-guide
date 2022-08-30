package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type logWriter struct{}

func main() {

	resp, err := http.Get("http://google.com")

	if err != nil {
		// log the error and quit the program
		log.Fatal("Error: ", err)
	}

	// Make a slice and initailize it with 99999 spaces to accomodate the response
	// bs := make([]byte, 99999)
	
	// Response struct has the body property which implements the reader interface
	// Can do the read function
	// resp.Body.Read(bs)
	// Returns the HTML of the page
	// fmt.Println(string(bs))

	lw := logWriter{}

	// resp.body returns a byte slice pointer 
	// Copy takes information from within our application and writes it out
	// Body = implements reader interface
	// Copy = implements writer interface (type File has a function called Write)
	io.Copy(lw, resp.Body)

}


func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}