package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

type Address struct {
	Street  string
	City    string
	Country string
}

func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	d := gob.NewDecoder(&b)
	result := Person{}
	_ = d.Decode(&result)
	return &result
}

func main() {
	adarsh := &Person{
		Name: "Adarsh",
		Address: &Address{
			Street:  "Road",
			City:    "Bonglao",
			Country: "India",
		},
		Friends: []string{"l1", "l2"},
	}

	ankit := adarsh.DeepCopy()

	ankit.Name = "ankit"
	ankit.Address.City = "gaon"
	ankit.Friends = append(ankit.Friends, "l32")

	fmt.Println(ankit, ankit.Address)
	fmt.Println(adarsh, adarsh.Address)

}
