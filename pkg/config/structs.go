package config

type Response struct {
	Artists   string
	Locations string
	Dates     string
	Relation  string
}

type Artist struct {
	Id          int
	Image       string
	Name        string
	CrDate      int    `json:"creationDate"`
	FstAlbum    string `json:"firstAlbum"`
	Members     []string
	RelLink     string `json:"Relations"`
	Relations   RelationsData
	Coordinates [][2]float64
}

type RelationsData struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Concerts struct {
	Index []struct {
		Id        int
		Locations []string
		Dates     string
	}
}

type ArtistSe struct {
	Param string
	Type  string
	Link  int
}

type Result struct {
	Artist  []ArtistSe
	Band    []ArtistSe
	Member  []ArtistSe
	EDate   []ArtistSe
	FDate   []ArtistSe
	Concert []ArtistSe
}

type ResultG struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		}
	}
}
