package handlers

import (
	"encoding/json"
	"fmt"
	"groupieFunc/pkg/config"
	"groupieFunc/pkg/functions"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var GOOGLE_API_KEY = "GOOGLE_API_KEY"
var artists []config.Artist
var concerts config.Concerts
var locations []string

// Set the variables(artists, concerts) from the main
func SetData(ar *[]config.Artist, co *config.Concerts, loc *[]string) {
	artists = *ar
	concerts = *co
	locations = *loc
}

// PageData
type PageData struct {
	Artists   []config.Artist
	Locations []string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Parse home HTML
	t, err := template.ParseFiles("./static/templates/home.html", "./static/templates/artists.html", "./static/templates/base.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Prepare the pagedata
	pageData := PageData{
		Artists:   artists,
		Locations: locations,
	}

	// Serve and Execute HTML
	if r.URL.Path == "/style.css" {
		http.ServeFile(w, r, "./static/css/styles.css")
		return
	} else if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	} else {
		t.Execute(w, pageData)
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	// Get the artist id
	artistId := r.URL.Query().Get("id")

	// Redirect if id is empty
	if artistId == "" {
		http.Redirect(w, r, "/artist?id=1", http.StatusSeeOther)
		return
	}

	i, errAtoi := strconv.Atoi(artistId)
	if errAtoi != nil || i < 1 || i > 52 {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	artistInfo := artists[i-1]

	// Pull artist relations
	relRes := functions.GetResponse(artistInfo.RelLink)
	var relations config.RelationsData
	json.Unmarshal(relRes, &relations)

	// Assign the relations to the artist
	artistInfo.Relations = relations

	// Get geocoding
	var urlLocations []string
	for location := range artistInfo.Relations.DatesLocations {
		location = strings.ReplaceAll(location, "_", "%20")
		location = strings.ReplaceAll(location, "-", "%20")
		urlLocations = append(urlLocations, location)
	}

	// Channel to receive coordinates from goroutines
	coordChan := make(chan [2]float64, len(urlLocations))
	var wg sync.WaitGroup

	for _, urlLocation := range urlLocations {
		wg.Add(1)
		go fetchCoordinates(urlLocation, &wg, coordChan)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(coordChan)

	// Receive coordinates from channel
	var coordinates [][2]float64
	for coord := range coordChan {
		coordinates = append(coordinates, coord)
	}

	artistInfo.Coordinates = coordinates

	// Execute the artist HTML
	t, err := template.ParseFiles("./static/templates/artist.html", "./static/templates/base.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	t.Execute(w, artistInfo)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Get the query serach data
	searchParam := r.URL.Query().Get("q")

	// Filter the artist based on the search param
	var result config.Result
	if len(searchParam) >= 1 {
		result = functions.GetResult(artists, concerts, searchParam)
	}

	// Parse the search template
	t, _ := template.ParseFiles("static/templates/search.html")
	t.Execute(w, result)
}

func ResultsHandler(w http.ResponseWriter, r *http.Request) {
	// Get the query serach data
	searchParam := r.URL.Query().Get("q")

	// Filter the artist based on the search param
	var result config.Result
	if len(searchParam) >= 1 {
		result = functions.GetResult(artists, concerts, searchParam)
	}

	t, err := template.ParseFiles("./static/templates/results.html", "./static/templates/base.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	t.Execute(w, result)
}

func FilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//Get the form values
		filterParam := map[string]string{}
		r.ParseForm()
		for key, value := range r.Form {
			if key[0] == 'n' {
				filterParam["nofmembers"] += value[0] + ","
			} else if len(value) > 0 && value[0] != "" {
				if key == "crDate" {
					filterParam["crDateStart"] = value[0][:4]
					filterParam["crDateLast"] = value[0][5:9]
				} else if key == "fstAlbumDate" {
					filterParam["fstAlbumStart"] = value[0][:4]
					filterParam["fstAlbumLast"] = value[0][5:9]
				} else {
					filterParam[key] = value[0]
				}
			}
		}
		log.Println(filterParam, r.Form)

		//Get the filtered result
		result := functions.GetFilter(artists, filterParam)

		//Parse the template and executed with the filtered result
		t, _ := template.ParseFiles("./static/templates/artists.html")
		t.Execute(w, result)
	}
}

func fetchCoordinates(urlLocation string, wg *sync.WaitGroup, coordChan chan<- [2]float64) {
	defer wg.Done()

	geoURL := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", urlLocation, GOOGLE_API_KEY)

	geoRes := functions.GetResponse(geoURL)

	var result config.ResultG
	err := json.Unmarshal(geoRes, &result)
	if err != nil {
		log.Fatal("Error unmarshalling JSON:", err)
	}

	if len(result.Results) > 0 {
		latitude := result.Results[0].Geometry.Location.Lat
		longitude := result.Results[0].Geometry.Location.Lng
		coord := [2]float64{latitude, longitude}
		coordChan <- coord
	}
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		t, err := template.ParseFiles("./static/templates/404.html", "./static/templates/base.html")
		if err != nil {
			fmt.Fprint(w, "404: Page not found")
			return
		}
		t.Execute(w, nil)
	} else if status == http.StatusInternalServerError {
		t, err := template.ParseFiles("./static/templates/500.html", "./static/templates/base.html")
		if err != nil {
			fmt.Fprint(w, "500: Internal Server Error")
			return
		}
		t.Execute(w, nil)
	}
}
