package main

import (
	"fmt"
	"net/http"
)

func checkLink(l string, c chan string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Println(l, " might be down!")
		c <- l

	} else {

		fmt.Println(l + " is fine")
		c <- l

	}

}

func main() {

	greeting := "hello wolf"

	myChan := make(chan string)

	go (func(g string, c chan string) {
		fmt.Println(g)
		c <- "done"

	})(greeting, myChan)

	fmt.Println(<-myChan)

	// links := []string{
	// 	"http://amazon.com",
	// 	"http://arsebook.com",
	// 	"http://google.com",
	// 	"http://stackoverflow.com",
	// 	"http://golang.org",
	// }

	// c := make(chan string)

	// for _, link := range links {
	// 	go checkLink(link, c)
	// }

	// for link := range c {
	// 	go func(linky string) { //function literal
	// 		time.Sleep(5 * time.Second)
	// 		checkLink(linky, c)
	// 	}(link)
	// }

}
