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
}
