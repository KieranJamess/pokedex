package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationData struct {
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
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
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

func (c *Client) GetLocationData(areaName string) (LocationData, error) {
	endpoint := "/location-area/" + areaName
	url := BASE_URL + endpoint
	var locationData LocationData

	fmt.Printf("Checking cache for: %s - ", url)
	data, ok := c.cache.Get(url)
	if ok {
		fmt.Println("Found in cache")
		locationData = LocationData{}
		err := json.Unmarshal(data, &locationData)
		if err != nil {
			return locationData, err
		}

		return locationData, nil
	}
	fmt.Println("Not found in cache")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationData, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return locationData, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return locationData, fmt.Errorf(
			"request failed with status code %d",
			res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return locationData, err
	}

	locationData = LocationData{}
	err = json.Unmarshal(data, &locationData)
	if err != nil {
		return locationData, err
	}

	c.cache.Add(url, data)

	return locationData, nil
}
