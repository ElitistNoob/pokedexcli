package api

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func (c *Client) GetLocations(endpoint *string) (Locations, error) {
	target := BaseUrl + LocationArea
	if endpoint != nil {
		target = *endpoint
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

	return locations, nil
}
