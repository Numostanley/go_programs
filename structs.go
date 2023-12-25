package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	sex       string
}

func (p person) fullName() string {
	return fmt.Sprintf(
		"%v %v",
		p.firstName,
		p.lastName,
	)
}

func main() {

	numo := person{
		"Numo",
		"Stanley",
		27,
		"Male",
	}
	fmt.Println("My name is", numo.fullName())
}
