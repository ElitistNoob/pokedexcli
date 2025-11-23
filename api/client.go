package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	cache "github.com/ElitistNoob/pokedexcli/internal/pokecache"
)

type Client struct {
	cache  cache.Cache
	client http.Client
}

func NewRequest(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		cache: cache.NewCache(cacheInterval),
		client: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) MakeRequest(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	return c.client.Do(req)
}

func (c *Client) fetchJSON(target string, T any) error {
	if cachedData, ok := c.cache.Get(target); ok {
		return json.Unmarshal(cachedData, T)
	}

	res, err := c.MakeRequest("GET", target, nil)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if err := json.NewDecoder(res.Body).Decode(T); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	data, err := json.Marshal(T)
	if err == nil {
		c.cache.Add(target, data)
	}

	return nil
}
