package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/tomocy/meander"
)

func main() {
	http.HandleFunc("/recommendations", withCORS(handleRecommendations))
	http.ListenAndServe(":8080", nil)
}

func handleRecommendations(w http.ResponseWriter, r *http.Request) {
	urlQuery := r.URL.Query()
	q := meander.Query{
		Radius:  1000,
		Journey: strings.Split(urlQuery.Get("journey"), "|"),
	}
	q.Lat, _ = strconv.ParseFloat(urlQuery.Get("lat"), 64)
	q.Lng, _ = strconv.ParseFloat(urlQuery.Get("lng"), 64)
	q.Radius, _ = strconv.Atoi(urlQuery.Get("radius"))
	q.CostRangeStr = urlQuery.Get("cost")
	places := q.Run()
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
