package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func (englishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Hi There!"
}
func (spanishBot) getGreeting() string {
	// Very custom logic for generating an spanish greeting
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
