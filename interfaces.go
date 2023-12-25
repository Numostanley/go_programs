package main

import "fmt"

type car interface {
	getModel() string
	getYear() string
}

type benz struct {
	car // letting the type know the interface it satisfies

	year  string
	model string
	size  string
}

func (c benz) getModel() string {
	return c.model
}

func (c benz) getYear() string {
	return c.year
}

func main() {

	myCar := benz{
		year:  "2023",
		model: "GLE450",
		size:  "medium",
	}

	var newCar car

	newCar = myCar

	fmt.Println("Model:", newCar.getModel())
	fmt.Println("Year", newCar.getYear())

}
