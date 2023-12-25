package main

import "fmt"

type car interface {
	getModel() string
	getYear() string
}

type benz struct {
	year string
	model string
	size string
}

func (c benz) getModel() string {
	return c.model
}

func (c benz) getYear() string {
	return c.year
}

type toyota struct {
	year string
	model string
	fuelGuage string
}

func (c toyota) getModel() string {
	return c.model
}

func (c toyota) getYear() string {
	return c.year
}


// interface assertion
func getFuelGuage(c car) string {
	e, ok := c.(toyota)

	if ok {
		return e.fuelGuage
	}

	return "No fuel guage"
}

// type switching
func getFullSPec(c car) string {
	switch e := c.(type) {
	case benz:
		return fmt.Sprintf("%s %s %s", e.year, e.model, e.size,)

	case toyota:
		return fmt.Sprintf("%s %s %s", e.year, e.model, e.fuelGuage,)

	default:
		return ""
	}

}

func main() {

	myCar := benz {
		year: "2023",
		model: "GLE450",
		size: "medium",
	}

	ourCar := toyota {
		year: "2023",
		model: "GLE450",
		fuelGuage: "50litres",
	}

	var newCar car

	newCar = myCar

	fmt.Println("Model:", newCar.getModel())
	fmt.Println("Year", newCar.getYear())
	fmt.Println("Fuel Guage - Toyota:", getFuelGuage(ourCar))
	fmt.Println("Fuel Guage - Benz:", getFuelGuage(newCar))
	fmt.Println("Full Spec - Benz:", getFullSPec(newCar))
	fmt.Println("Full Spec - Toyota:", getFullSPec(ourCar))

}
