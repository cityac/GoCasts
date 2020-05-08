package main

import "fmt"

type englishBot struct {
	greeting string
}
type spanishBot struct {
	greeting string
}

type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{greeting: "Hello"}
	sb := spanishBot{greeting: "Halo"}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// getGreeting
// englishBot imlements bot interface implicitly now
func (eb englishBot) getGreeting() string {
	return eb.greeting + "!"
}

// getGreeting
// spanishBot imlements bot interface implicitly now
func (sb spanishBot) getGreeting() string {
	return sb.greeting + "!"
}
