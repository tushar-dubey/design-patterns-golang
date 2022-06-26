package main

import "fmt"

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

func (a *Address) DeepCopy() *Address {
	return &Address{
		Street:  a.Street,
		City:    a.City,
		Country: a.Country,
	}
}

func (p *Person) DeepCopy() *Person {
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
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
