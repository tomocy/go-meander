package meander

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Query struct {
	Lat          float64
	Lng          float64
	Journey      []string
	Radius       int
	CostRangeStr string
}

func (q Query) Find(types string) (*googleResponse, error) {
	endpoint := "https://maps.googleapis.com/maps/api/place/nearbysearch/json"
	vals := make(url.Values)
	vals.Set("location", fmt.Sprintf("%g,%g", q.Lat, q.Lng))
	vals.Set("radius", fmt.Sprintf("%d", q.Radius))
	vals.Set("types", types)
	vals.Set("key", APIKey)
	if 0 < len(q.CostRangeStr) {
		costRange := ParseCostRange(q.CostRangeStr)
		vals.Set("minprice", fmt.Sprintf("%d", int(costRange.From)-1))
		vals.Set("maxprice", fmt.Sprintf("%d", int(costRange.To)-1))
	}

	resp, err := http.Get(endpoint + vals.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var googleResp googleResponse
	if err := json.NewDecoder(resp.Body).Decode(&googleResp); err != nil {
		return nil, err
	}

	return &googleResp, nil
}

func (q Query) Run() []interface{} {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	places := make([]interface{}, len(q.Journey))
	for i, j := range q.Journey {
		wg.Add(1)
		go func(types string, i int) {
			defer wg.Done()
			resp, err := q.Find(types)
			if err != nil {
				log.Printf("could not search places: %s\n", err)
				return
			}
			if len(resp.Results) == 0 {
				log.Println("no results")
				return
			}

			n := rand.Intn(len(resp.Results))
			result := resp.Results[n]
			for _, photo := range result.Photos {
				photo.URL = "https://maps.googleapis.com/maps/api/place/photo?maxwith=1000" + "&photoreference=" + photo.PhotoRef + "&key=" + APIKey
			}
			places[i] = result
		}(j, i)
	}

	wg.Wait()
	return places
}
