package main

import "fmt"

// Renderer We want to render circle triangle and square with Vector and Raster rendering
type Renderer interface {
	RenderCircle(radius float32)
	RenderEquilateralTriangle(side int)
}

type VectorRenderer struct {
	//
}

func (v *VectorRenderer) RenderEquilateralTriangle(side int) {
	fmt.Println("Drawing an Equilateral triangle with side: ", side)
}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of Radius: ", radius)
}

type RasterRenderer struct {
	//
}

func (r *RasterRenderer) RenderEquilateralTriangle(side int) {
	fmt.Println("Rendering an Equilateral triangle with side: ", side)
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

type EquilateralTriangle struct {
	side     int
	renderer Renderer
}

func NewEquilateralTriangle(side int, renderer Renderer) *EquilateralTriangle {
	return &EquilateralTriangle{side: side, renderer: renderer}
}

func (e *EquilateralTriangle) Draw() {
	e.renderer.RenderEquilateralTriangle(e.side)
}

func main() {
	rasterRenderer := RasterRenderer{}
	vectorRenderer := VectorRenderer{}

	circle := NewCircle(&rasterRenderer, 10)
	circle.Draw()

	triangle := NewEquilateralTriangle(15, &vectorRenderer)
	triangle.Draw()
}
