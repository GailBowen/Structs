package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("Please pass in a file as an argument")
		os.Exit(1)
	}

	fmt.Println(os.Args)

	myFile, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, myFile)

	//bytesFromFile, readError := os.ReadFile(os.Args[1])

	// if readError != nil {
	// 	fmt.Println("Could not read from file")
	// 	os.Exit(1)
	// }
	// stringFromBytes := string(bytesFromFile) //Convert byte slice to string
	// fmt.Println(stringFromBytes)

}
