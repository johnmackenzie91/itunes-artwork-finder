package finder

import (
	"context"

	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/env"
	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/pkg/itunes"

	"github.com/sirupsen/logrus"
)

type Func func(ctx context.Context, term, country, entity string) (SearchResponse, error)

func Itunes(e env.Config) Func {
	itunesCli, _ := itunes.New(itunes.SetDomain(e.ItunesEndpoint), itunes.WithLogger(logrus.New())) //TODO: handle error
	return Func(func(ctx context.Context, term, country, entity string) (SearchResponse, error) {
		res, err := itunesCli.Search(ctx, term, country, entity)

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
	})
}
