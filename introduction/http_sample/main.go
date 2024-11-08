package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Home Page")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler)
	router.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the about page")
	})
	http.ListenAndServe(":8080", nil)
}
