package main

import "fmt"

type englishBot struct {
	name     string
	language string
}

func (eb englishBot) getGreeting() string {
	return "My name is " + eb.name
}

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
