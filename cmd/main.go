package main

import (
	"encoding/json"
	"net/http"

	"github.com/tomocy/meander"
)

func main() {
	http.HandleFunc("/recommendations", withCORS(handleRecommendations))
	http.ListenAndServe(":8080", nil)
}

func handleRecommendations(w http.ResponseWriter, r *http.Request) {
	query := meander.NewPlaceSearchQuery(r.URL.Query())
	places := query.Run()
	respond(w, places)
}

func respond(w http.ResponseWriter, data []interface{}) error {
	publicData := make([]interface{}, len(data))
	for i, d := range data {
		publicData[i] = meander.Public(d)
	}
	return json.NewEncoder(w).Encode(publicData)
}

func withCORS(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		f(w, r)
	}
}
