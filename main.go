package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/3")

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)

}
