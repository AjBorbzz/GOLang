package main

import (
	"fmt"
	"strings"
)

// Crete Book structure
type Book struct {
	Title  string
	Author string
}

type Library struct {
	Books []Book
}

func (l *Library) AddBook(title, author string) {
	l.Books = append(l.Books, Book{Title: title, Author: author})
	fmt.Println("Books added successfully.")
}

func (l *Library) ViewBooks() {
	if len(l.Books) == 0 {
		fmt.Println("No books available")
	}
	fmt.Println("Books in the library.")
	for i, book := range l.Books {
		fmt.Printf("%d. %s by %s \n", i+1, book.Title, book.Author)
	}
}

func (l *Library) DeleteBook(title string) {
	for i, book := range l.Books {
		if strings.ToLower(book.Title) == strings.ToLower(title) {
			l.Books = append(l.Books[:i], l.Books[i+1:]...) // Use to remove the book
			fmt.Println("Book deleted successfully.")
			return
		}
	}
	fmt.Println("Book not found.")
}

func main() {
	var library Library
	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add a book")
		fmt.Println("2. View all books")
		fmt.Println("3. Delete a book")
		fmt.Println("4. Exit")
		fmt.Print("Please choose an option (1-4): ")
	}
}
