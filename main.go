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
