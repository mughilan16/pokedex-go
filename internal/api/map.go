package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type locationList struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

const (
	locationBaseUrl = "https://pokeapi.co/api/v2/location?offset=%d&limit=%d"
	LocationLimit   = 20
)

func FetchLocation(offset int) (locationList, error) {
	url := fmt.Sprintf(locationBaseUrl, offset, LocationLimit)
	res, err := http.Get(url)
	if err != nil {
		return locationList{}, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return locationList{}, err
	}
	return parseLocation(data)
}

func parseLocation(data []byte) (locationList, error) {
	list := locationList{}
	err := json.Unmarshal(data, &list)
	if err != nil {
		return locationList{}, err
	}
	return list, nil
}
