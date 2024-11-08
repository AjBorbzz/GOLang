package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// PageData represents the data to be passed into the HTML template
type PageData struct {
	Title   string
	Heading string
	Message string
}

// HomePageHandler renders the home page with dynamic data
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	// Data that will be passed to the template
	data := PageData{
		Title:   "My Go Website with Gorilla Mux",
		Heading: "Welcome to My Home Page",
		Message: "This page is dynamically rendered using Go templates and Gorilla Mux!",
	}

	// Parse the HTML template file
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Execute the template with the provided data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func main() {
	// Initialize Gorilla Mux router
	r := mux.NewRouter()

	// Define the route for the homepage
	r.HandleFunc("/", HomePageHandler)

	// Start the web server
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
