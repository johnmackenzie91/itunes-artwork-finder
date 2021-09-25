package artwork

import (
	"context"

	itunesCli "bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/pkg/itunes"
)

// Itunes returns an adapter to the itunes service
func itunes(client itunesCli.Client) Adapter {
	return Adapter(func(ctx context.Context, term, country, entity string) (SearchResponse, error) {
		res, err := client.Search(ctx, term, country, entity)

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
