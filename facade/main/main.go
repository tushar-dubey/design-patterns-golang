package main

import "fmt"

type Buffer struct {
	Width, Height int
	buffer        []rune
}

func NewBuffer(width int, height int) *Buffer {
	return &Buffer{Width: width, Height: height, buffer: make([]rune, width*height)}
}

func (b *Buffer) At(index int) rune {
	return b.buffer[index]
}

type ViewPort struct {
	buffer *Buffer
	offset int
}

func NewViewPort(buffer *Buffer) *ViewPort {
	return &ViewPort{buffer: buffer}
}

func (v *ViewPort) GetCharacterAt(index int) rune {
	return v.buffer.At(v.offset + index)
}

type Console struct {
	Buffers   []*Buffer
	Viewports []*ViewPort
	offset    int
}

func NewConsole() *Console {
	b := NewBuffer(200, 150)
	v := NewViewPort(b)
	return &Console{Buffers: []*Buffer{b}, Viewports: []*ViewPort{v}, offset: 0}
}

func (c *Console) GetCharacterAt(index int) rune {
	return c.Viewports[0].GetCharacterAt(index)
}

func main() {

	// Console is a facade on top of Buffer and Viewports but they both are still available for use directly
	c := NewConsole()
	fmt.Println(c.GetCharacterAt(0))
}
