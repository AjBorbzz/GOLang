package main

import (
	"fmt"
	"os"
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

		var option int
		_, err := fmt.Scan(&option)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 4.")
			continue
		}

		// Handle different cases using switch-case
		switch option {
		case 1:
			// Add a book
			var title, author string
			fmt.Print("Enter the book title: ")
			fmt.Scan(&title)
			fmt.Print("Enter the author's name: ")
			fmt.Scan(&author)
			library.AddBook(title, author)

		case 2:
			// View all books
			library.ViewBooks()

		case 3:
			// Delete a book
			var title string
			fmt.Print("Enter the title of the book to delete: ")
			fmt.Scan(&title)
			library.DeleteBook(title)

		case 4:
			// Exit the program
			fmt.Println("Exiting Library Management System...")
			os.Exit(0)

		default:
			// Invalid option
			fmt.Println("Invalid option. Please choose a valid option between 1 and 4.")
		}
	}
}
