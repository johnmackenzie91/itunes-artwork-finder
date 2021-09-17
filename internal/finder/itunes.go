package finder

import (
	"context"
	"fmt"

	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/env"
	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/pkg/itunes"
)

type Wrapper struct {
	itunesClient itunes.Client
}

// New takes our env struct and attempts to initialize a client
func New(e env.Config) (Wrapper, error) {
	itunesCli, err := itunes.New(itunes.SetDomain(e.ItunesEndpoint)) //TODO: add logger
	if err != nil {
		return Wrapper{}, fmt.Errorf("failed to initialize client: %w", err)
	}
	return Wrapper{
		itunesClient: itunesCli,
	}, nil
}

type SearchResponse itunes.SearchResponse

func (w Wrapper) Search(ctx context.Context, term, country, entity string) (SearchResponse, error) {
	res, err := w.itunesClient.Search(ctx, term, country, entity)

	if err != nil {
		return SearchResponse{}, err
	}

	return SearchResponse(res), nil
}
