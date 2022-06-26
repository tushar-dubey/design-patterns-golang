package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
)

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "Developer", 110000}
	case Manager:
		return &Employee{"", "Manager", 150000}
	default:
		panic("unsupported type bro")
	}
}

func main() {
	dev := NewEmployee(Developer)
	dev.Name = "Test"
	fmt.Println(dev)
}
