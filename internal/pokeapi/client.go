package pokeapi

import (
	"net/http"
	"time"

	"github.com/mrmaths1212/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	pokeCache  pokecache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout time.Duration, interval time.Duration) Client {
	return Client{
		pokeCache: pokecache.newCache(interval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
