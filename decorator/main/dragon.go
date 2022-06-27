package main

import "fmt"

type Aged interface {
	Age() int
	SetAge(int)
}

type Bird struct {
	age int
}

func (b *Bird) Age() int {
	return b.age
}

func (b *Bird) SetAge(age int) {
	b.age = age
}

func (b *Bird) Fly() {
	if b.Age() > 10 {
		fmt.Println("Flying!")
	}
}

type Lizard struct {
	age int
}

func (l *Lizard) Age() int {
	return l.age
}

func (l *Lizard) SetAge(age int) {
	l.age = age
}

func (l *Lizard) Crawl() {
	if l.Age() < 10 {
		fmt.Println("Crawling!")
	}
}

type Dragon struct {
	bird   Bird
	lizard Lizard
}

func (d *Dragon) SetAge(age int) {
	d.bird.SetAge(age)
	d.lizard.SetAge(age)
}

func (d *Dragon) Age() int {
	return d.bird.Age()
}

func (d *Dragon) Fly() {
	d.bird.Fly()
}

func (d *Dragon) Crawl() {
	d.lizard.Crawl()
}

func NewDragon() *Dragon {
	return &Dragon{
		bird:   Bird{},
		lizard: Lizard{},
	}
}
