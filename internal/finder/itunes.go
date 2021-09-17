package finder

import (
	"context"
	"fmt"

	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/env"
	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/pkg/itunes"

	"github.com/sirupsen/logrus"
)

type Wrapper struct {
	itunesClient itunes.Client
}

// New takes our env struct and attempts to initialize a client
func New(e env.Config) (Wrapper, error) {
	itunesCli, err := itunes.New(itunes.SetDomain(e.ItunesEndpoint), itunes.WithLogger(logrus.New())) //TODO: updated logger logic
	if err != nil {
		return Wrapper{}, fmt.Errorf("failed to initialize client: %w", err)
	}
	return Wrapper{
		itunesClient: itunesCli,
	}, nil
}

type SearchResponse struct {
	ResultCount uint `json:"resultCount"`
	Results     []Result
}

type Result struct {
	Artist         string
	CollectionName string
	Link           string
}

func (w Wrapper) Search(ctx context.Context, term, country, entity string) (SearchResponse, error) {
	res, err := w.itunesClient.Search(ctx, term, country, entity)

	if err != nil {
		return SearchResponse{}, err
	}

	out := SearchResponse{}
	out.ResultCount = res.ResultCount
	for _, row := range res.Results {
		r := Result{Artist: row.ArtistName, CollectionName: row.CollectionName, Link: row.ArtworkURL100}
		out.Results = append(out.Results, r)
	}

	return out, nil
}
