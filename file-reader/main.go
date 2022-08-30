package main

import (
	"io"
	"log"
	"os"
)

func main() {
	arg := os.Args[1]

	// Returns a pointer to a file and an error
	file, err := os.Open(arg) // For read access.
	
	if err != nil {
		log.Fatal("Error: ", err)
	}

	// File implements reader interface.  We can use Copy to print out the content
	io.Copy(os.Stdout,file)

	// Make a byte slice to hold the content of the file
	// data := make([]byte, 100)

	// Read the content of the file into the byte slice created
	// file.Read(data)

	// fmt.Println((string(data)))
}