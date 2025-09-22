package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + name

	if val, ok := c.cache.Get(url); ok {
		cacheVal := RespPokemon{}
		err := json.Unmarshal(val, &cacheVal)
		if err != nil {
			return RespPokemon{}, nil
		}

		return cacheVal, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, nil
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, nil
	}

	c.cache.Add(url, data)

	pokemon := RespPokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return RespPokemon{}, nil
	}

	return pokemon, nil
}
