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

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p *person) print() {
	fmt.Printf("%+v", *p)
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// using shortcut pointer
	// same as &jim.updateName("Jimmy")
	jim.updateName("Jimmy")
	jim.print()
}

func SampleAlex() {
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contactInfo = contactInfo{email: "alex@gmail.com", zipCode: 214}
	alex = person{firstName: "Alex", lastName: "Anderson"}
	alex.print()
	fmt.Printf("%+v", alex)

}
