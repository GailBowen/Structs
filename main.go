package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

func main() {
	fmt.Println("hello structs")

	bob := person{firstname: "Gail", lastname: "Foad"}

	// fmt.Println(bob)

	var Romanov person
	Romanov.firstname = "Nicolas"

	fmt.Println(Romanov.firstname)
	fmt.Printf("%+v", bob)

}
