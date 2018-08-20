package meander

import "os"

var APIKey string

func init() {
	APIKey = os.Getenv("GOOGLE_PLACES_API_KEY")
}
