package main

import "fmt"

type Vehicle struct {
	Make  string
	Model string
	Year  int
}

func (v Vehicle) Details() string {
	return fmt.Sprintf("The Car's spec: %d %s %s", v.Year, v.Make, v.Model)
}

type Car struct {
	Vehicle
	Doors int
}

func (c Car) CarDetails() string {
	return fmt.Sprintf("%s with %d doors", c.Details(), c.Doors)
}

func main() {
	myCar := Car{
		Vehicle: Vehicle{
			Make:  "Toyota",
			Model: "Corolla Altis",
			Year:  2022,
		},
		Doors: 4,
	}

	fmt.Println(myCar.CarDetails())
}
