package main

import (
	"fmt"
	"groupieFunc/pkg/handlers"
	"net/http"
)

func main() {
	// Hello msg
	fmt.Print("\nHi, go to http://localhost:8000/ to view the site!\n")

	// Handlers
	http.HandleFunc("/", handlers.Handler)
	http.HandleFunc("/artist", handlers.ArtistHandler)
	http.HandleFunc("/search", handlers.SearchHandler)

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Port init
	http.ListenAndServe(":8000", nil)
}
