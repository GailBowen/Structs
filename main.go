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
	i := "gail"
	fmt.Printf("1. main  -- i  %T: &i=%p i=%v\n", i, &i, i)
	namePointer := &i
	byVal(namePointer)
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
