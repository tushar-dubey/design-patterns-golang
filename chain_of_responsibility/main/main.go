package main

import "fmt"

func main() {
	// Creature example with modifiers for chain of responsibility
	goblin := &Creature{
		Name:    "goblin",
		Attack:  1,
		Defense: 1,
	}
	fmt.Println(goblin.String())
	root := NewCreatureModifier(goblin)
	root.Add(NewDoubleAttackModifier(goblin))
	root.Add(NewDoubleDefenseModifier(goblin))
	root.Handle()
	fmt.Println(goblin.String())
}
