package functions

import (
	"io"
	"log"
	"net/http"
)

func GetResponse(url string) []byte {
	// API request to pull the data
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Error while fetching the data", err)
	}
	defer res.Body.Close()

	// Parse the JSON
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error while reading the body", err)
	}

	return body
}
