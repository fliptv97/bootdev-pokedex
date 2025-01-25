package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationAreaResponse struct {
	Count   int     `json:"count"`
	Next    *string `json:"next"`
	Prev    *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

const baseUrl = "https://pokeapi.co/api/v2"

func (c *Client) LoadLocationAreas(pageUrl *string) (LocationAreaResponse, error) {
	url := baseUrl + "/location-area"

	if pageUrl != nil {
		url = *pageUrl
	}

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

	var locationAreaRes LocationAreaResponse

	if err := json.Unmarshal(rawData, &locationAreaRes); err != nil {
		return LocationAreaResponse{}, err
	}

	return locationAreaRes, nil
}
