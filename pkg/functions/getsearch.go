package functions

import (
	"fmt"
	"groupieFunc/pkg/config"
	"strconv"
	"strings"
)

func GetResult(artists []config.Artist, concerts config.Concerts, searchParam string) []config.ArtistSe {
	var result []config.ArtistSe
	searchParam = strings.ToLower(searchParam)
	for _, artist := range artists {
		// Search artist/band
		if strings.Contains(strings.ToLower(artist.Name), searchParam) {
			var artistSe config.ArtistSe
			artistSe.Param = artist.Name
			if len(artist.Members) > 1 {
				artistSe.Type = "Band"
			} else {
				artistSe.Type = "Artist"
			}
			artistSe.Link = artist.Id
			result = append(result, artistSe)
		}

		// Search member
		if len(artist.Members) != 1 {
			for _, member := range artist.Members {
				if strings.Contains(strings.ToLower(member), searchParam) {
					var artistSe config.ArtistSe
					artistSe.Param = fmt.Sprintf("%s - %s", artist.Name, member)
					artistSe.Type = "Member"
					artistSe.Link = artist.Id
					result = append(result, artistSe)
				}
			}
		}

		// Search dates
		if strings.Contains(artist.FstAlbum, searchParam) {
			var artistSe config.ArtistSe
			artistSe.Param = artist.FstAlbum
			artistSe.Type = fmt.Sprintf("First Album: %s", artist.Name)
			artistSe.Link = artist.Id
			result = append(result, artistSe)
		}
		crDate := strconv.Itoa(artist.CrDate)
		if strings.Contains(crDate, searchParam) {
			var artistSe config.ArtistSe
			artistSe.Param = artist.FstAlbum
			artistSe.Type = fmt.Sprintf("Established: %s", artist.Name)
			artistSe.Link = artist.Id
			result = append(result, artistSe)
		}
	}

	// Search concerts
	for _, artistStruct := range concerts.Index {
		for _, artistConcerts := range artistStruct.Locations {
			if strings.Contains(artistConcerts, searchParam) {
				var artistSe config.ArtistSe
				artistSe.Param = artistConcerts
				artistSe.Type = "Concert"
				artistSe.Link = artistStruct.Id
				result = append(result, artistSe)
			}
		}
	}

	return result
}
