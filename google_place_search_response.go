package meander

type googlePlacesAPIResponse struct {
	Places []*googlePlace `json:"results"`
	Status string         `json:"status"`
}

type googlePlace struct {
	Name     string `json:"name"`
	IconURL  string `json:"icon"`
	Vicinity string `json:"vicinity"`
	Geometry struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
	} `json:"geometry"`
	Photos []struct {
		Reference string `json:"photo_reference"`
		URL       string `json:"url"`
	} `json:"photos"`
}

func (p *googlePlace) setPhotoURLs() {
	for i := 0; i < len(p.Photos); i++ {
		p.Photos[i].URL = "https://maps.googleapis.com/maps/api/place/photo?maxwith=1000&photoreference=" + p.Photos[i].Reference + "&key=" + googlePlacesAPIKey
	}
}
