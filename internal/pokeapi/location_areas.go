package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreas, error) {
	endpoint := "/location-area/"
	url := BASE_URL + endpoint
	var locationAreas LocationAreas

	if pageURL != nil {
		url = *pageURL
	}

	fmt.Printf("Checking cache for: %s - ", url)
	data, ok := c.cache.Get(url)
	if ok {
		fmt.Println("Found in cache")
		locationAreas = LocationAreas{}
		err := json.Unmarshal(data, &locationAreas)
		if err != nil {
			return locationAreas, err
		}

		return locationAreas, nil
	}
	fmt.Println("Not found in cache")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationAreas, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreas, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return locationAreas, fmt.Errorf(
			"request failed with status code %d",
			res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return locationAreas, err
	}

	locationAreas = LocationAreas{}
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return locationAreas, err
	}

	c.cache.Add(url, data)

	return locationAreas, nil
}
