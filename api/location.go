package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	cache "github.com/ElitistNoob/pokedexcli/internal/pokecache"
)

type Locations struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (c *Client) GetLocations(endpoint *string, cache *cache.Cache) (Locations, error) {
	target := BaseUrl + LocationArea
	if endpoint != nil {
		target = *endpoint
	}

	cachedData, isCached := cache.Get(target)
	if isCached {
		var locations Locations
		if err := json.Unmarshal(cachedData, &locations); err != nil {
			return Locations{}, fmt.Errorf("failed to Unmarshall cached data: %w", err)
		}
		fmt.Println("Data loaded from Cache!")
		fmt.Println("====================================")
		return locations, nil
	}

	res, err := c.MakeRequest("GET", target, nil)
	if err != nil {
		return Locations{}, fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Locations{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	var locations Locations
	if err := json.NewDecoder(res.Body).Decode(&locations); err != nil {
		return Locations{}, fmt.Errorf("failed to decode response: %w", err)
	}

	data, err := json.Marshal(locations)
	if err != nil {
		fmt.Printf("Warning: Failed to Marshal locations: %v\n", err)
	} else {
		cache.Add(target, data)
	}

	fmt.Println("Data loaded from the World Wide Web!")
	fmt.Println("====================================")
	return locations, nil
}
