package main

import "fmt"

type Creature struct {
	Name            string
	Attack, Defense int
}

func (c *Creature) String() string {
	return fmt.Sprintf("Name: %s, (%d/%d)", c.Name, c.Attack, c.Defense)
}

type Modifier interface {
	Add(m Modifier)
	Handle()
}

type CreatureModifier struct {
	c    *Creature
	next Modifier
}

func (c *CreatureModifier) Add(m Modifier) {
	if c.next != nil {
		c.next.Add(m)
	} else {
		c.next = m
	}
}

func (c *CreatureModifier) Handle() {
	if c.next != nil {
		c.next.Handle()
	}
}

func NewCreatureModifier(c *Creature) *CreatureModifier {
	return &CreatureModifier{c: c}
}

type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifier(creature *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{CreatureModifier: CreatureModifier{c: creature}}
}

func (d *DoubleAttackModifier) Handle() {
	fmt.Println("Doubling ", d.c.Name, "'s Attack")
	d.c.Attack *= 2
	d.CreatureModifier.Handle()
}

type DoubleDefenseModifier struct {
	CreatureModifier
}

func NewDoubleDefenseModifier(creature *Creature) *DoubleDefenseModifier {
	return &DoubleDefenseModifier{CreatureModifier: CreatureModifier{
		c:    creature,
		next: nil,
	}}
}

func (d *DoubleDefenseModifier) Handle() {
	if d.c.Attack <= 2 {
		fmt.Println("Adding Defense to %s", d.c.Name)
		d.c.Defense++
	} else {
		fmt.Println("Attak is out of range")
	}
	d.CreatureModifier.Handle()
}
