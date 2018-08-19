package meander

type place struct {
	Name            string         `json:"name"`
	Icon            string         `json:"icon"`
	Vicinity        string         `json:"vicinity"`
	Photos          []*googlePhoto `json:"photos"`
	*googleGeometry `json:"geometry"`
}

type googlePhoto struct {
	PhotoRef string `json:"photo_reference"`
	URL      string `json:"url"`
}

type googleGeometry struct {
	*googleLocation `json:"location"`
}

type googleLocation struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
