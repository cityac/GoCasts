package main

import "fmt"

type contactInfo struct {
	email   string
	phone   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

// func main() {
// 	var alex person
// 	alex.firstName = "Alex"
// 	alex.lastName = "Anderson"
// 	fmt.Println(alex)
// 	fmt.Printf("%+v", alex)
// }

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			phone:   "+12163924794",
			zipCode: 80032,
		},
	}
	jim.updateName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
