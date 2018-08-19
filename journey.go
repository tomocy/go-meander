package meander

import "strings"

// Journeys is a slice of an arrary of interfaces which represents several journeys
var Journeys = []interface{}{
	&j{
		Name: "Romantic",
		PlaceTypes: []string{
			"park", "bar", "movie_theater", "restaurant",
			"florist", "taxi_stand",
		},
	},
	&j{
		Name: "Shopping",
		PlaceTypes: []string{
			"department_store", "cafe", "clothing_store", "jewelry_store",
			"shoe_store",
		},
	},
	&j{
		Name: "Night Life",
		PlaceTypes: []string{
			"bar", "casino", "food", "bar",
			"night_club", "bar", "bar", "hospital",
		},
	},
	&j{
		Name: "Culture",
		PlaceTypes: []string{
			"museum", "cafe", "cemetery", "library",
			"art_gallery",
		},
	},
	&j{
		Name: "Relax",
		PlaceTypes: []string{
			"hair_care", "beauty_salon", "cafe", "spa",
		},
	},
}

type j struct {
	Name       string
	PlaceTypes []string
}

func (j j) Public() interface{} {
	return map[string]interface{}{
		"name":     j.Name,
		"journeys": strings.Join(j.PlaceTypes, "|"),
	}
}
