package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseUrl + "/pokemon/" + strings.ToLower(name)

	_, ok := c.cache.Get(url)

	if !ok {
		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			return Pokemon{}, err
		}

		res, err := c.httpClient.Do(req)

		if err != nil {
			return Pokemon{}, err
		}

		defer res.Body.Close()

		rawData, err := io.ReadAll(res.Body)

		if err != nil {
			return Pokemon{}, err
		}

		c.cache.Add(url, rawData)
	}

	cachedData, _ := c.cache.Get(url)

	var pokemonRes Pokemon

	if err := json.Unmarshal(cachedData, &pokemonRes); err != nil {
		return Pokemon{}, err
	}

	return pokemonRes, nil
}
