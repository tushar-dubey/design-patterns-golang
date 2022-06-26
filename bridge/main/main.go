package main

import "fmt"

// Renderer We want to render circle triangle and square with Vector and Raster rendering
type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRenderer struct {
	//
}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of Radius: ", radius)
}

type RasterRenderer struct {
	//
}

func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Rendering a circle of Radius: ", radius)
}

type Circle struct {
	renderer Renderer
	radius   float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func main() {

}
