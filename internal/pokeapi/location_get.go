package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetLocation retrieves location area data for the given location name.
// It checks the cache first, and if not found, performs an HTTP GET request.
func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	// Try to get data from cache
	if val, ok := c.cache.Get(url); ok {
		locationResp := Location{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return Location{}, err
		}
		return locationResp, nil
	}

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	// Execute the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	// Read response body
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	// Parse JSON into Location struct
	locationResp := Location{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return Location{}, err
	}

	// Store the data in cache for future use
	c.cache.Add(url, dat)
	return locationResp, nil
}
