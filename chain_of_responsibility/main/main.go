package main

import (
	"fmt"
	"sync"
)

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

	// Second Creature example with COR CQS Observer Mediator

	game := &Game{observers: sync.Map{}}
	secondGoblin := NewSecondCreature(game, "Strong Goblin", 2, 2)
	fmt.Println(secondGoblin.String())

	{
		m := NewDoubleAttackSecondModifier(game, secondGoblin)
		fmt.Println(secondGoblin.String())
		_ = m.Close()
	}
	fmt.Println(secondGoblin.String())

}
