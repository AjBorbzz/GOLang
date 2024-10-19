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

/*

1. Vehicle Struct:
   - The `Vehicle` struct has fields for `Make`, `Model`, and `Year`.
   - The `Details()` method returns a formatted string representing the vehicle's details.

2. Car Struct:
   - The `Car` struct embeds the `Vehicle` struct, which means it inherits its fields and methods.
   - The `Car` struct also has an additional field, `Doors`, specific to cars.
   - The `CarDetails()` method combines the vehicle details with the number of doors.

3. Main Function:
   - An instance of `Car` is created, with the embedded `Vehicle` initialized with its fields.
   - The `CarDetails()` method is called to print the full details of the car.

Use Case

This pattern of struct composition is commonly used in Go for several reasons:

- Code Reusability: You can define shared behaviors (methods) in the parent struct and reuse them in child structs.
- Logical Grouping: It helps in logically grouping related types, making your code cleaner and easier to maintain.
- Polymorphism: It allows you to use a single interface to handle different types that share a common structure or behavior.

*/
