package location

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"sort"
	"sync"
	"time"

	"github.com/Antony15/goodWorkLabs-Test/constants"
	"github.com/Antony15/goodWorkLabs-Test/logger"
	"github.com/Antony15/goodWorkLabs-Test/mapResponse"
	"github.com/Antony15/goodWorkLabs-Test/redis"
)

// Global variable to initialize redis client
var redisClient = redis.Initialize()

// Request struct
type request struct {
	Latitude  *string `json:"latitude"`
	Longitude *string `json:"longitude"`
}

// Return Response struct
type response struct {
	Results results `json:"results"`
}

// Results struct
type results struct {
	Items []items `json:"items"`
}

// Items struct
type items struct {
	Type    string    `json:"type"`
	Closest []closest `json:"closest_POIs"`
}

// Closest struct
type closest struct {
	Distance      int     `json:"distance"`
	Title         string  `json:"title"`
	AverageRating float64 `json:"averageRating"`
	ID            string  `json:"id"`
}

// New function used as constructor to return pointer of request
func New() *request {
	return &request{}
}

// ValidateRequest method receiver used to validate the incoming request required fields
func (req *request) ValidateRequest() error {
	if req.Latitude == nil || req.Longitude == nil {
		return errors.New("Error : Lattitude or Longitude should not be null")
	}
	return nil
}

// SendRequest method receiver used to send request to Here Map Api & find 3 closest POI’s of each type
func (req *request) SendRequest() (response, bool) {
	var (
		response response
		result   results
	)
	// type of places
	places := []string{"Parking spots", "Charging Stations", "Restaurants"}
	wg := sync.WaitGroup{}
	items := make([]items, len(places))
	// loop over places & find closest POI's
	for k, v := range places {
		// redis key
		key := *req.Latitude + "," + *req.Longitude + v
		// create url params
		params := url.Values{}
		params.Add("at", *req.Latitude+","+*req.Longitude)
		params.Add("q", v)
		params.Add("apiKey", constants.HMapApiKey)
		// endpoint url
		url := constants.HMapReqEP + params.Encode()
		wg.Add(1)
		// concurrent function to send request to each type
		go func(url string, i int, val string, rediskey string) {
			mp := mapResponse.New()
			// get cached memory from redis if available
			if err := redisClient.GetKey(rediskey, mp); err != nil {
				// if not key available, make a get request to Here map api
				resp, err := http.Get(url)
				checkErr(err)
				err = json.NewDecoder(resp.Body).Decode(&mp)
				// save returned response from Here map api in redis cache
				if err = redisClient.SetKey(rediskey, &mp, time.Minute*constants.RedisExpInMin); err != nil {
					logger.Log.Println("Error in setting redis key: ", rediskey, err.Error())
				}
			}
			closest := make([]closest, len(mp.Results.Items))
			for k1, v1 := range mp.Results.Items {
				closest[k1].Distance = v1.Distance
				closest[k1].AverageRating = v1.AverageRating
				closest[k1].ID = v1.ID
				closest[k1].Title = v1.Title
			}
			items[i].Type = val
			items[i].Closest = findClosestPoi(closest)
			wg.Done()
		}(url, k, v, key)
	}
	wg.Wait()
	result.Items = items
	response.Results = result
	return response, true
}

// findClosestPoi function is used to find the closest POI's of each type
func findClosestPoi(closest []closest) []closest {
	sort.Slice(closest, func(i, j int) bool {
		return closest[i].Distance < closest[j].Distance
	})
	if len(closest) <= 3 {
		return closest
	}
	return closest[:3]
}

// checkErr function is used to check for errors & log in log file
func checkErr(e error) {
	if e != nil {
		logger.Log.Println("Error : ", e.Error())
	}
}
