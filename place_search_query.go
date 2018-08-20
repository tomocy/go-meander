package meander

import (
	"net/url"
)

type PlaceSearchQuery interface {
	Run() []interface{}
}

func NewPlaceSearchQuery(vals url.Values) PlaceSearchQuery {
	return newGooglePlaceSearchQuery(vals)
}
