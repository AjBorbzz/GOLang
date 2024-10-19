package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func NewPerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func (p Person) Display() string {
	return fmt.Sprintf("Name: %s %s, Age: %d", p.FirstName, p.LastName, p.Age)
}

func main() {
	person := NewPerson("Aj", "Borbe", 37)
	fmt.Println(person.Display())
}
