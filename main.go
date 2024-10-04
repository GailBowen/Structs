package main

import (
	"fmt"
)

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

	colours := map[string]string{
		"purple": "#800080",
		"green":  "#00FF00",
		"brown":  "#964B00",
	}

	printMap(colours)

	fmt.Printf("%v+", colours)

	sl := []int{3, 7, 2}

	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}

}

func printMap(c map[string]string) {

	for colour, hex := range c {
		fmt.Println(colour, hex)
	}

	c["silver"] = "#C0C0C0"
}

func byVal(q *string) {
	fmt.Printf("2. byVal -- q %T: &q=%p q=&i=%p  *q=i=%v\n", q, &q, q, *q)
}

func updateSlice(s []string) {
	s[0] = "Bonjour"
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateLastName(newLastName string) {
	(*p).lastname = newLastName
}
