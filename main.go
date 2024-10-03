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

	// var colours map[string]string

	// colours := map[string]string{
	// 	"purple": "#800080",
	// 	"green":  "#00FF00",
	// }

	//fmt.Println(colours)

	colours := make(map[string]string)
	colours["purple"] = "#800080"
	colours["green"] = "#00FF00"

	fmt.Printf("%+v\n", colours)

	delete(colours, "purple")

	fmt.Printf("%+v\n", colours)

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
