package pokeapi

import (
	"log"
	"net/http"
	"time"

	"github.com/samkc-0/repl-template/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	pokecache  pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	cache, err := pokecache.NewCache(5 * time.Second)
	if err != nil {
		log.Fatal(err)
	}
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokecache: cache,
	}
}
