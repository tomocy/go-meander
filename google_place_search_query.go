package meander

import (
	"bufio"
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

type googlePlaceSearchQuery struct {
	lat          float64
	lng          float64
	journeys     []string
	radius       int
	costRangeStr string
}

func newGooglePlaceSearchQuery(vals url.Values) *googlePlaceSearchQuery {
	q := new(googlePlaceSearchQuery)
	q.journeys = strings.Split(vals.Get("journey"), "|")
	q.lat, _ = strconv.ParseFloat(vals.Get("lat"), 64)
	q.lng, _ = strconv.ParseFloat(vals.Get("lng"), 64)
	q.radius, _ = strconv.Atoi(vals.Get("radius"))
	q.costRangeStr = vals.Get("cost")

	return q
}

func (q googlePlaceSearchQuery) Run() []interface{} {
	var wg sync.WaitGroup
	placesCh := make(chan interface{}, len(q.journeys))
	for _, journey := range q.journeys {
		wg.Add(1)
		go q.findAndDeliverPlaceRandomly(placesCh, journey, wg.Done)
	}
	wg.Wait()
	close(placesCh)

	return receivePlaces(placesCh)
}

func (q googlePlaceSearchQuery) findAndDeliverPlaceRandomly(placeCh chan<- interface{}, journey string, deferF func()) {
	defer deferF()
	resp, err := q.find(journey)
	if err != nil {
		log.Printf("could not search places: %s\n", err)
		return
	}
	if len(resp.Results) == 0 {
		log.Println("no results, may be rate limit")
		return
	}

	placeCh <- pickResultRandomly(resp.Results)
}

func (q googlePlaceSearchQuery) find(journey string) (*googleResponse, error) {
	vals := q.prepareURLValuesForGooglePlaceSearch(journey)
	resp, err := makeRequestForGooglePlaceSearch(vals.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	log.Println(resp.Status)

	var googleResp googleResponse
	if err := json.NewDecoder(resp.Body).Decode(&googleResp); err != nil {
		if err != io.EOF {
			return nil, err
		}
	}

	// for _, result := range googleResp.Results {
	// 	result.setPhotoURLs()
	// }

	return &googleResp, nil
}

func (q googlePlaceSearchQuery) prepareURLValuesForGooglePlaceSearch(journy string) url.Values {
	vals := make(url.Values)
	vals.Set("location", fmt.Sprintf("%g,%g", q.lat, q.lng))
	vals.Set("radius", fmt.Sprintf("%d", q.radius))
	vals.Set("type", journy)
	vals.Set("key", APIKey)
	if 0 < len(q.costRangeStr) {
		costRange := parseCostRange(q.costRangeStr)
		vals.Set("minprice", fmt.Sprintf("%d", int(costRange.from)-1))
		vals.Set("maxprice", fmt.Sprintf("%d", int(costRange.to)-1))
	}

	return vals
}

func makeRequestForGooglePlaceSearch(encodedVals string) (*http.Response, error) {
	endpoint := "https://maps.googleapis.com/maps/api/place/nearbysearch/json"
	resp, err := http.Get(endpoint + "?" + encodedVals)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func decode(r io.Reader, data interface{}) error {
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		if err != io.EOF {
			return err
		}
	}
	return nil
}

func pickResultRandomly(results []*place) *place {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(results))
	return results[n]
}

func receivePlaces(placesCh <-chan interface{}) []interface{} {
	places := make([]interface{}, 0)
	for place := range placesCh {
		places = append(places, place)
	}

	return places
}
