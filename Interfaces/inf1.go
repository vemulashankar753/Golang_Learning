package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type Bot interface {
	getGreeting() string
}

func main() {
	fmt.Println("Hello World")
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func (eb englishBot) getGreeting() string {
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b Bot) string {
	value := b.getGreeting()
	fmt.Println(value)
	return value
}

//The below one gives the error as "golang" does not support the Method
//overloading, we cannot have 2 functions with same name, but different args

//func printGreeting(eb englishBot) string {
//	return eb.getGreeting()
//}

//func printGreeting(sb spanishBot) string {
//	return
//}
//
//shankar
