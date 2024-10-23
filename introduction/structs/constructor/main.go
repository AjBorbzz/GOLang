package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func NewBook(title string, author string, pages int) Book {
	return Book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
}

func main() {
	book1 := NewBook("1984", "George Orwell", 328)
	book2 := NewBook("The Catcher in the Rye", "J.D Salinger", 277)

	fmt.Println("Book 1: %+v\n", book1)
	fmt.Println("Book 2: %+v\n", book2)
}
