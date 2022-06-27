package main

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle has radius %f ", c.Radius)
}

func (c *Circle) Resize(r float32) {
	c.Radius *= r
}

type Square struct {
	Side float32
}

func (s Square) Render() string {
	return fmt.Sprintf("Square has side %f ", s.Side)
}

type ColoredShape struct {
	Shape
	color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s of color %s ", c.Shape.Render(), c.color)
}

type TransparentShape struct {
	Shape
	transparency float32
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%s has transparency %f", t.Shape.Render(), t.transparency*100)
}
