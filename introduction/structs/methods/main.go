package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greeting() string {
	return fmt.Sprintf("Hello %s, your age is : %d", p.Name, p.Age)
}

func main() {
	p := Person{
		Name: "Aj Borbe",
		Age:  37,
	}
	fmt.Println(p.Greeting())
}
