package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	exploreBaseUrl = "https://pokeapi.co/api/v2/location-area/%s"
)

type exploreData struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func FetchExplore(args []string) (exploreData, error) {
	if len(args) < 1 {
		// return exploreData{}, errors.New("Not enough arguments")
		return exploreData{}, errors.New("")
	}
	areaName := args[0]
	areaName = "pastoria-city-area"
	url := fmt.Sprintf(exploreBaseUrl, areaName)
	// fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		return exploreData{}, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return exploreData{}, err
	}
	return parseExplore(data)
}

func parseExplore(data []byte) (exploreData, error) {
	list := exploreData{}
	err := json.Unmarshal(data, &list)
	if err != nil {
		return exploreData{}, err
	}
	// fmt.Println(list)
	return list, nil
}
