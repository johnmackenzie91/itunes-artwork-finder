package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/app/middleware/logging"
	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/app/redoc"
	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/domain"
	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/finder"

	"github.com/go-chi/chi"
	"github.com/johnmackenzie91/commonlogger"
)

var _ ServerInterface = (*handlers)(nil)

//go:generate mockery --name Searcher
type Searcher interface {
	Search(ctx context.Context, term, country, entity string) (finder.SearchResponse, error)
}

// New implements the implmentation of the interface generated from the openapi spec.
func New(client Searcher, logger commonlogger.ErrorInfoDebugger) http.Handler {
	r := chi.NewMux()

	// init request/response middleware
	r.Use(logging.LoggingMiddleware(logger))

	// init the documentation endpoints
	docEndpoints := redoc.New(logger)
	r.Get("/docs", docEndpoints.V1Docs)
	r.Get("/docs/spec", docEndpoints.V1Spec)

	// init status endpoint for health check
	r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK")) //nolint: errcheck #nothing be be gained from this check
	})
	return HandlerFromMux(handlers{Client: client, logger: logger}, r)
}

type handlers struct {
	Client Searcher
	logger commonlogger.ErrorInfoDebugger
}

func (s handlers) GetRestV1AlbumSearch(w http.ResponseWriter, r *http.Request, params GetRestV1AlbumSearchParams) {
	w.Header().Add("Content-Type", "application/json")

	if err := params.Validate(); err != nil {
		s.writeErrorResponse(r.Context(), w, WrapError(err, ErrBadRequest))
		return
	}

	// search itunes
	searchTerm := params.Title
	if params.Artist != nil {
		searchTerm = fmt.Sprintf("%s - %s", *params.Artist, searchTerm)
	}
	out, err := s.Client.Search(r.Context(), searchTerm, "gb", "album")

	if err != nil {
		s.writeErrorResponse(r.Context(), w, WrapError(err, ErrBadGateway))
		return
	}

	rtn := []domain.Album{}

	for _, item := range out.Results {
		artistMatch := true
		if params.Artist != nil {
			artistMatch = normalise(item.Artist) == *params.Artist
		}

		titlesMatch := strings.Contains(normalise(item.CollectionName), params.Title)

		if artistMatch == false || titlesMatch == false {
			s.logger.Debug("dropping result", item)
			continue
		}

		row := domain.Album{
			ImageURL:   item.Link,
			Title:      item.CollectionName,
			ArtistName: item.Artist,
		}

		if params.Size != nil {
			// Update return itunes return values to match the size requested by OUR client
			row.ImageURL = strings.Replace(row.ImageURL, "100x100", fmt.Sprintf("%sx%s", params.Size, params.Size), 1)
		}
		rtn = append(rtn, row)
	}

	if len(rtn) == 0 {
		s.writeErrorResponse(r.Context(), w, ErrNotFound)
		return
	}

	// write response
	b, err := json.Marshal(rtn)
	if err != nil {
		s.writeErrorResponse(r.Context(), w, WrapError(err, ErrInternal))
		return
	}
	if _, err = w.Write(b); err != nil {
		s.writeErrorResponse(r.Context(), w, WrapError(err, ErrInternal))
		return
	}
}

// normalise removes unwanted chars from client input
func normalise(input string) string {
	input = strings.ToLower(input)
	input = strings.Replace(input, "+", " ", -1)
	return input
}

func (s handlers) writeErrorResponse(ctx context.Context, w http.ResponseWriter, err error) {
	var out sentinelAPIError

	s.logger.Error(ctx, err)

	switch e := err.(type) {
	case sentinelWrappedError:
		out = e.sentinel
	case sentinelAPIError:
		out = e
	case error:
		out.Code = http.StatusInternalServerError
		out.Msg = "internal server error"
	}

	// write response code
	w.WriteHeader(out.Code)

	// write error to output in json format
	if _, err := w.Write(out.JSON()); err != nil {
		panic(err)
	}
}
