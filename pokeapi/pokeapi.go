package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

type Locations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocations(url *string) (Locations, error) {
	if url == nil {
		url = new(string)
		*url = "https://pokeapi.co/api/v2/location/"
	}

	req, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		return Locations{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}
	defer res.Body.Close()

	var locs Locations
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locs)
	if err != nil {
		return Locations{}, err
	}
	return locs, nil
}

func Hello() {
	fmt.Println("hello world")
}
