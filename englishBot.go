package main

type englishBot struct {
	name     string
	language string
}

func (eb englishBot) getGreeting() string {
	return "My name is " + eb.name
}
