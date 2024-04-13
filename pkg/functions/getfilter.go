package functions

import (
	"groupieFunc/pkg/config"
	"log"
)

func GetFilter(artists []config.Artist) []config.Artist {
	log.Println("start filtering")
	return artists
}
