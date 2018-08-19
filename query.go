package meander

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Query struct {
	Lat          float64
	Lng          float64
	Journeys     []string
	Radius       int
	CostRangeStr string
}

func NewQuery(vals url.Values) *Query {
	q := new(Query)
	q.Journeys = strings.Split(vals.Get("journey"), "|")
	q.Lat, _ = strconv.ParseFloat(vals.Get("lat"), 64)
	q.Lng, _ = strconv.ParseFloat(vals.Get("lng"), 64)
	q.Radius, _ = strconv.Atoi(vals.Get("radius"))
	q.CostRangeStr = vals.Get("cost")

	return q
}

func (q Query) Find(types string) (*googleResponse, error) {
	endpoint := "https://maps.googleapis.com/maps/api/place/nearbysearch/json"
	vals := make(url.Values)
	vals.Set("location", fmt.Sprintf("%g,%g", q.Lat, q.Lng))
	vals.Set("radius", fmt.Sprintf("%d", q.Radius))
	vals.Set("type", types)
	vals.Set("key", APIKey)
	if 0 < len(q.CostRangeStr) {
		costRange := ParseCostRange(q.CostRangeStr)
		vals.Set("minprice", fmt.Sprintf("%d", int(costRange.From)-1))
		vals.Set("maxprice", fmt.Sprintf("%d", int(costRange.To)-1))
	}

	resp, err := http.Get(endpoint + "?" + vals.Encode())
	if err != nil {
		log.Println("query could not http.Get")
		return nil, err
	}
	defer resp.Body.Close()
	log.Println(resp.Status)

	var googleResp googleResponse
	if err := json.NewDecoder(resp.Body).Decode(&googleResp); err != nil {
		if err != io.EOF {
			return nil, err
		}
	}

	return &googleResp, nil
}

func (q Query) Run() []interface{} {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	places := make([]interface{}, len(q.Journeys))
	for i, journey := range q.Journeys {
		wg.Add(1)
		go func(journey string, i int) {
			defer wg.Done()
			resp, err := q.Find(journey)
			if err != nil {
				log.Printf("could not search places: %s\n", err)
				places = append(places[:i], places[i+1:])
				return
			}
			if len(resp.Results) == 0 {
				log.Println("no results")
				places = append(places[:i], places[i+1:])
				return
			}

			places[i] = pickResultRandomly(resp.Results)
		}(journey, i)
	}

	wg.Wait()
	return places
}

func pickResultRandomly(results []*place) *place {
	n := rand.Intn(len(results))
	return results[n]
}
