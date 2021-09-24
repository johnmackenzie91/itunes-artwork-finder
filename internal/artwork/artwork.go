package artwork

import (
	"context"

	itunesCli "bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/pkg/itunes"

	"github.com/johnmackenzie91/commonlogger"
)

// Adapter is what will be returned from this package
type Adapter func(ctx context.Context, term, country, entity string) (SearchResponse, error)

// Itunes is a constructor func that creates the itunes artwork adapter
func Itunes(endpoint string, logger commonlogger.ErrorInfoDebugger) (Adapter, error) {
	cli, err := itunesCli.New(
		itunesCli.SetDomain(endpoint),
		itunesCli.WithLogger(logger),
	)
	if err != nil {
		return nil, err
	}
	return itunes(cli), nil
}
