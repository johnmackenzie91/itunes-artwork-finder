package client

import (
	"net/http"
	"time"

	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/env"

	"github.com/gregjones/httpcache"
)

// New creates a new client with a cached transport layer
func New(e env.Config) *http.Client {
	return &http.Client{
		Timeout:   time.Duration(e.HTTPClientTimeout) * time.Second,
		Transport: httpcache.NewMemoryCacheTransport(),
	}
}
