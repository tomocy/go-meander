package meander

import "os"

var googlePlacesAPIKey string

func init() {
	googlePlacesAPIKey = os.Getenv("GOOGLE_PLACES_API_KEY")
}
