package functions

import (
	"fmt"
	"groupieFunc/pkg/config"
	"strconv"
	"strings"
)

func GetResult(artists []config.Artist, concerts config.Concerts, searchParam string) config.Result {
	var result config.Result
	searchParam = strings.ToLower(searchParam)
	for _, artist := range artists {
		// Search artist/band
		if strings.Contains(strings.ToLower(artist.Name), searchParam) {
			var artistSe config.ArtistSe
			artistSe.Param = artist.Name
			artistSe.Link = artist.Id
			if len(artist.Members) > 1 {
				artistSe.Type = "Band"
				result.Band = append(result.Band, artistSe)
			} else {
				artistSe.Type = "Artist"
				result.Artist = append(result.Artist, artistSe)
			}
		}

		// Search member
		if len(artist.Members) != 1 {
			for _, member := range artist.Members {
				if strings.Contains(strings.ToLower(member), searchParam) {
					var artistSe config.ArtistSe
					artistSe.Param = fmt.Sprintf("%s - %s", member, artist.Name)
					artistSe.Type = "Member"
					artistSe.Link = artist.Id
					result.Member = append(result.Member, artistSe)
				}
			}
		}

		// Search dates
		if strings.Contains(artist.FstAlbum, searchParam) {
			var artistSe config.ArtistSe
			artistSe.Param = fmt.Sprintf("%s - %s", artist.FstAlbum, artist.Name)
			artistSe.Type = "First Album"
			artistSe.Link = artist.Id
			result.FDate = append(result.FDate, artistSe)
		}
		crDate := strconv.Itoa(artist.CrDate)
		if strings.Contains(crDate, searchParam) {
			var artistSe config.ArtistSe
			artistSe.Param = fmt.Sprintf("%s - %s", crDate, artist.Name)
			artistSe.Type = "Established"
			artistSe.Link = artist.Id
			result.EDate = append(result.EDate, artistSe)
		}
	}

	// Search concerts
	for _, artistStruct := range concerts.Index {
		for _, artistConcerts := range artistStruct.Locations {
			if strings.Contains(artistConcerts, searchParam) {
				artist := artists[artistStruct.Id-1]
				var artistSe config.ArtistSe
				artistSe.Param = fmt.Sprintf("%s - %s", artistConcerts, artist.Name)
				artistSe.Type = "Concert"
				artistSe.Link = artistStruct.Id
				result.Concert = append(result.Concert, artistSe)
			}
		}
	}

	return result
}
