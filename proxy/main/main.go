package main

func main() {
	// example of protection proxy by using a safe car
	car := NewSafeCar(&Driver{17})
	car.Drive()

	// bitmap example of virtual proxy with lazy loading
	bitmap := NewBitmap("demo.png")
	DrawImage(bitmap)
	lazyMap := &LazyBitmap{filename: "demo2.png"}
	DrawImage(lazyMap)
}
