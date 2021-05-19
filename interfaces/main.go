package main

import "fmt"

type bot interface {
	getGreetings() string
}

type englishBot struct{}

type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreetings() string {

	// very custom logic for generating english greeting
	return "Hi, There!!"
}

func (spanishBot) getGreetings() string {
	return "Hola, Gringo!"
}

/* func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreetings())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeeting())
} */

func printGreeting(b bot) {
	fmt.Println(b.getGreetings())
}
