package main

import (
	"encoding/json"
	"fmt"
	"groupieFunc/pkg/config"
	"groupieFunc/pkg/functions"
	"groupieFunc/pkg/handlers"
	"net/http"
)

// Global variables
var GROUPIE_API = "https://groupietrackers.herokuapp.com/api"
var GROUPIE_LOCATIONS_API = "https://groupietrackers.herokuapp.com/api/locations"
var PORT = "8000"

func main() {
	// Pull the response data
	mainRes := functions.GetResponse(GROUPIE_API)
	var response config.Response // Response instance
	json.Unmarshal(mainRes, &response)

	// Pull artists data
	var artists []config.Artist
	artistRes := functions.GetResponse(response.Artists)
	json.Unmarshal([]byte(artistRes), &artists)

	// Pull locations(concerts)
	var concerts config.Concerts
	concertRes := functions.GetResponse(GROUPIE_LOCATIONS_API)
	json.Unmarshal(concertRes, &concerts)

	handlers.SetData(&artists, &concerts)

	// Hello msg
	fmt.Printf("\nHi, go to http://localhost:%s/ to view the site!\n", PORT)

	// Handlers
	http.HandleFunc("/", handlers.Handler)
	http.HandleFunc("/artist", handlers.ArtistHandler)
	http.HandleFunc("/search", handlers.SearchHandler)
	http.HandleFunc("/results", handlers.ResultsHandler)
	http.HandleFunc("/filter", handlers.FilterHandler)

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Port init
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil)
}
