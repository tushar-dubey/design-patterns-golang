package main

import "fmt"

//Example of Protection Proxy

type Driven interface {
	Drive()
}

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Car is being Driven!")
}

type Driver struct {
	Age int
}

type SafeCar struct {
	Car    Car
	Driver *Driver
}

func (s *SafeCar) Drive() {
	if s.Driver.Age >= 16 {
		fmt.Println("Car is being Driven!")
	} else {
		fmt.Println("Driver too young to drive")
	}
}

func NewSafeCar(driver *Driver) *SafeCar {
	return &SafeCar{Car: Car{}, Driver: driver}
}
