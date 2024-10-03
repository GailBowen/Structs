package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	contactInfo
}

type contactInfo struct {
	email    string
	houseNum int
}

func main() {
	fmt.Println("hello structs")

	bob := person{
		firstname: "Gail",
		lastname:  "Bowen",
		contactInfo: contactInfo{
			email:    "bob@gmail.com",
			houseNum: 76,
		},
	}

	bob.print()

	bob.updateLastName("Foad")

	bob.print()

}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateLastName(newLastName string) {
	(*p).lastname = newLastName
}
