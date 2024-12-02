package pokeapi

import (
	"github.com/KieranJamess/pokedex/internal/pokecache"
	"net/http"
	"time"
)

const (
	BASE_URL = "https://pokeapi.co/api/v2/"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute, //After a minute, kill the request.
		},
	}
}
