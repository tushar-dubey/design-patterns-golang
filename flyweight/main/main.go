package main

import "fmt"

func main() {
	// First using FormattedText which does not use flyweight pattern
	text := "This is an example of a text string"
	ft := NewFormattedText(text)
	ft.Capitalize(10, 17)
	fmt.Println(ft.String())

	// Using BetterFormattedText with Flyweight Pattern
	bft := NewBetterFormattedText(text)
	bft.Range(19, 20).Capital = true
	fmt.Println(bft.String())

	// Using a normal user struct having a name
	adarsh := NewUser("Adarsh Balak")
	fmt.Println(adarsh.FullName)

	// using a Flyweight User2 having a name
	nalayak := NewUser2("Nalayak Balak")
	fmt.Println(nalayak.FullName())
}
