package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetLocation gets the pokemon inside a location
func (c *Client) GetLocation(area string) (RespLocation, error) {
	url := baseURL + "/location-area/" + area

	if val, ok := c.cache.Get(url); ok {
		cacheVal := RespLocation{}
		err := json.Unmarshal(val, &cacheVal)
		if err != nil {
			return RespLocation{}, err
		}

		return cacheVal, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocation{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocation{}, nil
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocation{}, nil
	}

	c.cache.Add(url, dat)

	locationResp := RespLocation{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespLocation{}, nil
	}
	return locationResp, nil
}
