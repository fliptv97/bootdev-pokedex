package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) LoadLocationAreas(pageUrl *string) (ShallowLocationAreasResponse, error) {
	url := baseUrl + "/location-area"

	if pageUrl != nil {
		url = *pageUrl
	}

	_, ok := c.cache.Get(url)

	if !ok {
		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			return ShallowLocationAreasResponse{}, err
		}

		res, err := c.httpClient.Do(req)

		if err != nil {
			return ShallowLocationAreasResponse{}, err
		}

		defer res.Body.Close()

		rawData, err := io.ReadAll(res.Body)

		if err != nil {
			return ShallowLocationAreasResponse{}, err
		}

		c.cache.Add(url, rawData)
	}

	cachedData, _ := c.cache.Get(url)

	var locationAreasRes ShallowLocationAreasResponse

	if err := json.Unmarshal(cachedData, &locationAreasRes); err != nil {
		return ShallowLocationAreasResponse{}, err
	}

	return locationAreasRes, nil
}

func (c *Client) LoadLocationAreaByName(name string) (LocationAreaResponse, error) {
	url := baseUrl + "/location-area/" + name

	_, ok := c.cache.Get(url)

	if !ok {
		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			return LocationAreaResponse{}, err
		}

		res, err := c.httpClient.Do(req)

		if err != nil {
			return LocationAreaResponse{}, err
		}

		defer res.Body.Close()

		rawData, err := io.ReadAll(res.Body)

		if err != nil {
			return LocationAreaResponse{}, err
		}

		c.cache.Add(url, rawData)
	}

	cachedData, _ := c.cache.Get(url)

	var locationAreaRes LocationAreaResponse

	if err := json.Unmarshal(cachedData, &locationAreaRes); err != nil {
		return LocationAreaResponse{}, err
	}

	return locationAreaRes, nil
}
