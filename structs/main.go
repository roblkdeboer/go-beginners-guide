package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// Embed a struct within a struct
	contactInfo
}

func main() {	
	// Assigns a zero value field
	var james person
	fmt.Printf("%+v", james)
	james.firstName = "James"
	james.lastName = "Smith"
	fmt.Println(james)

	alex := person{
		firstName: "Alex", 
		lastName: "Anderson", 
		contactInfo: contactInfo{
			email: "alex.email.com", 
			zipCode: 234567,
		},
	}
	
	// & indicates the memory location of the variable alex
	// alexPointer := &alex

	alex.updateName("Alexander")
	alex.print()
}

// Change receiver type to *person.  A pointer that points at a person
// Go can automatically turn the type of person into the pointer for the person
func (pointerToPerson *person) updateName(newFirstName string) {
	// * gives the value that exists at the memory address
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}