package main

import (
	"encoding/json"
	"net/http"

	"github.com/tomocy/meander"
)

func main() {
	http.HandleFunc("/journeys", func(w http.ResponseWriter, r *http.Request) {
		respond(w, meander.Journeys)
	})
	http.ListenAndServe(":8080", nil)
}

func respond(w http.ResponseWriter, data []interface{}) error {
	return json.NewEncoder(w).Encode(data)
}
