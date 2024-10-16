package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{
		Name: "Aj Borbe",
		Age:  37,
	}
	fmt.Println(p.Name)
	fmt.Println(p.Age)
}
