package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 94000,
		},
	}
	alex.updateName("Alexy")
	alex.print()
}

//POINTERS: structs, int, float, bool, string
//NO POINTERS: slices, maps, channels, pointers, functions

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
