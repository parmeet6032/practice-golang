package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	// contactInfo // contactInfo contactInfo	-> shorthand
}

type contactInfo struct {
	email   string
	zipcode int
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}

func (p *person) updateName(firstName string, lastName string) {
	// p --> pointer to person
	(*p).firstName = firstName
	(*p).lastName = lastName
}

func main() {
	// // not preferred
	// alex := person{"Alex", "Anderson"}

	// preferred
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
	alex.print()

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipcode: 19873,
		},
	}

	// jimPointer := &jim
	// jimPointer.updateName("Jimmy", "Party")

	jim.updateName("Jimmy", "Party")

	jim.print()
}
