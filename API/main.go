package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

// Handler for the about route
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the About Page.")
}

// Handler for the contact route
func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contact us at: contact@example.com")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)

	port := ":8080"
	fmt.Println("Server is running on http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
