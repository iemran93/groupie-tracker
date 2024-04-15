package functions

import (
	"groupieFunc/pkg/config"
	"log"
	"strings"
)

func GetLocations(concerts config.Concerts) map[string][]string {
	locationsMap := map[string][]string{}
	locationsSet := map[string]bool{}
	for _, concert := range concerts.Index {
		for _, location := range concert.Locations {
			if _, ok := locationsSet[location]; !ok {
				parts := strings.Split(location, "-")
				country := parts[len(parts)-1]
				locationsMap[country] = append(locationsMap[country], location)
				locationsSet[location] = true
			}
		}
	}
	log.Println(locationsMap)
	return locationsMap
}
