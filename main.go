package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}

type contactInfo struct {
	email    string
	houseNum int
}

func main() {
	fmt.Println("hello structs")

	bob := person{
		firstname: "Gail",
		lastname:  "Foad",
		contact: contactInfo{
			email:    "bob@gmail.com",
			houseNum: 76,
		},
	}

	// fmt.Println(bob)

	var Romanov person
	Romanov.firstname = "Nicolas"

	fmt.Println(Romanov.firstname)
	fmt.Printf("%+v", bob)

}
