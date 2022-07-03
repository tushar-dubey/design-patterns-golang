package main

import (
	"fmt"
	"sync"
)

type Argument int

const (
	Attack Argument = iota
	Defense
)

type SecondCreature struct {
	game            *Game
	Name            string
	attack, defense int
}

func (s *SecondCreature) String() string {
	return fmt.Sprintf("%s (%d/%d)", s.Name, s.Attack(), s.Defense())
}

func NewSecondCreature(game *Game, name string, attack int, defense int) *SecondCreature {
	return &SecondCreature{game: game, Name: name, attack: attack, defense: defense}
}

func (s *SecondCreature) Attack() int {
	q := &Query{s.Name, Attack, s.attack}
	s.game.Fire(q)
	return q.Value
}

func (s *SecondCreature) Defense() int {
	q := &Query{s.Name, Defense, s.defense}
	s.game.Fire(q)
	return q.Value
}

// Mediator Pattern
type Game struct {
	observers sync.Map
}

func (g *Game) Subscribe(o Observer) {
	g.observers.Store(o, struct{}{})
}

func (g *Game) Unsubscribe(o Observer) {
	g.observers.Delete(o)
}

func (g *Game) Fire(q *Query) {
	g.observers.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		key.(Observer).Handle(q)
		return true
	})
}

type SecondCreatureModifier struct {
	game     *Game
	creature *SecondCreature
}

func (s *SecondCreatureModifier) Handle(q *Query) {
}

type DoubleAttackSecondModifier struct {
	SecondCreatureModifier
}

func (d *DoubleAttackSecondModifier) Close() error {
	d.game.Unsubscribe(d)
	return nil
}

func NewDoubleAttackSecondModifier(g *Game, c *SecondCreature) *DoubleAttackSecondModifier {
	d := &DoubleAttackSecondModifier{SecondCreatureModifier: SecondCreatureModifier{
		game:     g,
		creature: c,
	}}
	g.Subscribe(d)
	return d
}

func (d *DoubleAttackSecondModifier) Handle(q *Query) {
	if q.creatureName == d.creature.Name && q.WhatToQuery == Attack {
		q.Value *= 2
	}
}

// Observer Pattern
type Observer interface {
	Handle(q *Query)
}

type Observable interface {
	Subscribe(o Observer)
	Unsubscribe(o Observer)
	Fire(q *Query)
}

// Command Query Separation
type Query struct {
	creatureName string
	WhatToQuery  Argument
	Value        int
}
