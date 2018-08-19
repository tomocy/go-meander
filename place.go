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

func (p place) Public() interface{} {
	return map[string]interface{}{
		"name":     p.Name,
		"icon":     p.Icon,
		"vicinity": p.Vicinity,
		"photos":   p.Photos,
		"lat":      p.Lat,
		"lng":      p.Lng,
	}
}
