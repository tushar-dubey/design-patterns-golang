package main

func main() {
	// dragon example for decorator from bird and lizard
	dragon := NewDragon()
	dragon.SetAge(11)
	dragon.Fly()
	dragon.Crawl()
}
