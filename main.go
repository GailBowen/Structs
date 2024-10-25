package main

import (
	"net/http"
)

func checkLink(l string) (result string) {
	resp, err := http.Get(l)
	if err != nil {
		result := l + " has a problem\n"
		return result
	} else {

		result := l + " " + resp.Status + "\n"
		defer resp.Body.Close()

		return result
	}

}

func main() {
	links := []string{
		"http://google.com",
		"http://arsebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com"}

	for _, link := range links {
		println(checkLink(link))
	}

}
