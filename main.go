package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/2")

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	//fmt.Println(resp)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf(string(body))

}
