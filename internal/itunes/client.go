package itunes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/env"

	"github.com/johnmackenzie91/commonlogger"
)

const defaultDomain = "https://itunes.apple.com/search"

// Doer sends http requests
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is used for quering itunes endpoints
type Client struct {
	// Used to make the http requests.
	client Doer
	// The domain to send the requests to.
	domain string
	logger commonlogger.ErrorInfoDebugger
}

// NewClient implements a new itunes client
func NewClient(logger commonlogger.ErrorInfoDebugger, env env.Config) *Client {
	// Set the defaults.
	c := Client{
		client: http.DefaultClient,
		domain: defaultDomain,
		logger: logger,
	}

	if env.ItunesEndpoint != "" {
		c.domain = env.ItunesEndpoint
	}
	return &c
}

// Search searches the itunes api for the given parameters
func (c Client) Search(ctx context.Context, term, country, entity string) (SearchResponse, error) {
	r, err := c.buildSearchRequest(ctx, term, country, entity)

	if err != nil {
		return SearchResponse{}, fmt.Errorf("failed to build itunes search request: %w", err)
	}

	res, err := c.attemptRequest(r)

	if err != nil {
		return SearchResponse{}, fmt.Errorf("failed to search itunes: %w", err)
	}

	c.logger.Debug("response received", res)

	exp := SearchResponse{}
	if err := json.NewDecoder(res.Body).Decode(&exp); err != nil {
		return SearchResponse{}, err
	}

	return exp, res.Body.Close()
}
