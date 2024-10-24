package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	fmt.Println("hello wolf")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf(string(body))

}
