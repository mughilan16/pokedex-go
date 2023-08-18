package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	limit   = 20
	baseUrl = "https://pokeapi.co/api/v2/location?offset=%d&limit=%d"
)

func CommandMap() error {
	offset := 0
	list := fetch(offset)
	return nil
}

func fetch(offset int) locationList {
	url := fmt.Sprintf(baseUrl, offset, limit)
	res, err := http.Get(url)
	fmt.Println(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	list := locationList{}
	err = json.Unmarshal(body, &list)
	if err != nil {
		log.Fatalln(err)
	}
	return list
}
