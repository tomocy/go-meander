package meander

import "os"

var APIKey string

func init() {
	APIKey = os.Getenv("GOOGLE_PLACES_API_KEY")
}

type place struct {
	Name            string         `json:"name"`
	Icon            string         `json:"icon"`
	Vicinity        string         `json:"vicinity"`
	Photos          []*googlePhoto `json:"photos"`
	*googleGeometry `json:"geometry"`
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
