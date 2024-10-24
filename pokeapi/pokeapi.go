package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/Victor-AvestaAI/pokedex/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
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
		*url = "https://pokeapi.co/api/v2/location-area/"
	}

	if val, ok := c.cache.Get(*url); ok {
		locs := Locations{}
		err := json.Unmarshal(val, &locs)
		if err != nil {
			return Locations{}, err
		}
		return locs, nil
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

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Locations{}, nil
	}

	var locs Locations
	err = json.Unmarshal(data, &locs)
	if err != nil {
		return Locations{}, err
	}

	c.cache.Add(*url, data)

	return locs, nil
}
