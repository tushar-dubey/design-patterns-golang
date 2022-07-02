package main

import "fmt"

type Image interface {
	Draw()
}

type Bitmap struct {
	fileName string
}

func (b *Bitmap) Draw() {
	fmt.Println("Draw the Image from,", b.fileName)
}

func NewBitmap(fileName string) *Bitmap {
	fmt.Println("Loading image from file, ", fileName)
	return &Bitmap{fileName: fileName}
}

func DrawImage(image Image) {
	fmt.Println("About to draw the image")
	image.Draw()
	fmt.Println("Drew the image ")
}

type LazyBitmap struct {
	filename string
	Bitmap   *Bitmap
}

func (l *LazyBitmap) Draw() {
	if l.Bitmap == nil {
		l.Bitmap = NewBitmap(l.filename)
	}
	l.Bitmap.Draw()
}
