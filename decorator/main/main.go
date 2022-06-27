package main

import "fmt"

func main() {
	// dragon example for decorator from bird and lizard
	dragon := NewDragon()
	dragon.SetAge(11)
	dragon.Fly()
	dragon.Crawl()

	// shape example for decorator with coloring a circular shape and making it transparent
	c := Circle{Radius: 5}
	c.Resize(10)
	fmt.Println(c.Render())

	coloredCircle := ColoredShape{
		Shape: &c,
		color: "Blue",
	}
	fmt.Println(coloredCircle.Render())

	transparentColoredCircle := TransparentShape{
		Shape:        &coloredCircle,
		transparency: .6,
	}
	fmt.Println(transparentColoredCircle.Render())
}
