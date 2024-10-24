package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type logWriter struct {
	io.Writer
}

func (l logWriter) Write(p []byte) (n int, err error) {

	fmt.Printf("hello duck!\n\n")
	fmt.Printf(string(p) + "\n\n")

	return len(p), nil
}

func main() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/3")

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	bobWriter := logWriter{}

	mylen, err := io.Copy(bobWriter, resp.Body)

	fmt.Println(mylen)

}
