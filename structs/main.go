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
	
	alexPointer := &alex

	alexPointer.updateName("Alexander")
	alex.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}