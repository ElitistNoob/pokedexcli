package api

import (
	"io"
	"net/http"
	"time"
)

type Client struct {
	client http.Client
}

func NewRequest(timeout time.Duration) Client {
return Client {
		http.Client{
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
