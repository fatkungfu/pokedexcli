package pokeapi

import (
	"net/http"
	"time"

	"github.com/fatkungfu/pokedexcli/internal/pokecache"
)

// Client is a struct for making PokeAPI requests.
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient creates a new PokeAPI client with a specified timeout and cache interval.
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
