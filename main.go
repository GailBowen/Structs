package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

type germanBot struct {
	name     string
	language string
}

func (gb germanBot) getGreeting() string {
	return "My name is " + gb.name
}

func main() {
	engBot := englishBot{name: "Harry", language: "English"}
	gerBot := germanBot{name: "Kurt", language: "German"}

	printGreeting(engBot)
	printGreeting(gerBot)
}
