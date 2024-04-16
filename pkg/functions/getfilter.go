package functions

import (
	"groupieFunc/pkg/config"
	"log"
	"strconv"
	"strings"
)

var concerts config.Concerts

func SetConcerts(conc *config.Concerts) {
	concerts = *conc
}

func GetFilter(artists []config.Artist, filterPram map[string]string) []config.Artist {
	log.Println("start filtering")

	result := []config.Artist{}
	for _, artist := range artists {
		nOfMembers_pass := true
		crDateStart_pass := true
		crDateLast_pass := true
		fstAlbumStart_pass := true
		fstAlbumLast_pass := true
		location_pass := true

		//Filter by numbers of membres
		if nOfMembersStr, ok := filterPram["nofmembers"]; ok {
			nOfMembers, _ := strconv.Atoi(nOfMembersStr)
			if len(artist.Members) != nOfMembers {
				nOfMembers_pass = false
			}
		}

		//Filter by Creation Date
		if crDateStartStr, ok := filterPram["crDateStart"]; ok {
			crDateStart, _ := strconv.Atoi(crDateStartStr)
			if artist.CrDate < crDateStart {
				crDateStart_pass = false
			}
		}
		if crDateLastStr, ok := filterPram["crDateLast"]; ok {
			crDateLast, _ := strconv.Atoi(crDateLastStr)
			if artist.CrDate > crDateLast {
				crDateLast_pass = false
			}
		}

		//Filter by First Album Date
		artistFstAlbum, _ := strconv.Atoi(strings.Split(artist.FstAlbum, "-")[2])
		if fstAlbumStartStr, ok := filterPram["fstAlbumStart"]; ok {
			fstAlbumStart, _ := strconv.Atoi(fstAlbumStartStr)
			log.Println(artistFstAlbum)
			if artistFstAlbum < fstAlbumStart {
				crDateStart_pass = false
			}
		}
		if fstAlbumLastStr, ok := filterPram["fstAlbumLast"]; ok {
			fstAlbumLast, _ := strconv.Atoi(fstAlbumLastStr)
			if artistFstAlbum > fstAlbumLast {
				crDateStart_pass = false
			}
		}

		//Filter by location
		if location, ok := filterPram["locations"]; ok {
			found := false
			for _, loc := range concerts.Index[artist.Id-1].Locations {
				if location == loc {
					found = true
					break
				}
			}
			if !found {
				location_pass = false
			}
		}

		if nOfMembers_pass && crDateStart_pass && crDateLast_pass && fstAlbumStart_pass && fstAlbumLast_pass && location_pass {
			result = append(result, artist)
		}
	}
	return result
}
