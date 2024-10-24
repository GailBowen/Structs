package main

import (
	"fmt"
)

type bot interface { //Can't create a variable of type bot //interface type
	getGreeting() string
}

func doStuff(b bot) {
	fmt.Println(b.getGreeting())
}

type germanBot struct {
	name     string
	language string
}

func (gb germanBot) getGreeting() string { //Implementing this method is enough to make it type bot. //concrete type // Interfaces are implicit
	return "Ich heiße " + gb.name
}

type frenchBot struct {
}

func (fb frenchBot) getGreeting() string {
	return "Je m'appelle François"
}

func main() {
	engBot := englishBot{name: "Harry", language: "English"}
	gerBot := germanBot{name: "Kurt", language: "German"}
	frBot := frenchBot{}

	doStuff(frBot) //If you call a bot function, you have to implement all bot methods
	doStuff(engBot)
	doStuff(gerBot)
}
