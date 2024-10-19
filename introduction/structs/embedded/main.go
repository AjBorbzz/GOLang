package main

import "fmt"

type Address struct {
	Street string
	Number int
}

type Person struct {
	Name string
	Age  int
	Address
}

func main() {
	p := Person{
		Name: "Aj Borbe",
		Age:  37,
		Address: Address{
			Street: "Main Street",
			Number: 123,
		},
	}

	fmt.Println(p.Street)
	fmt.Println(p.Number)
}
