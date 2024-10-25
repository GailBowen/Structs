package main

import (
	"fmt"
	"net/http"
)

func checkLink(l string, c chan string) {
	resp, err := http.Get(l)
	if err != nil {
		fmt.Println(l + " has a problem")
		c <- "Might be down, I think\n"

	} else {

		fmt.Println(l + " " + resp.Status)
		c <- "Yes, it's up\n"

	}

}

func main() {
	links := []string{
		"http://amazon.com",
		"http://arsebook.com",
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c) //Receiving messages from a channel is a blocking line of code!
	}

}
