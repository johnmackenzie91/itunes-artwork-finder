package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/app/middleware/logging"
	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/domain"
	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/itunes"

	"github.com/go-chi/chi"
	"github.com/johnmackenzie91/commonlogger"
)

var _ ServerInterface = (*handlers)(nil)

// New implements the implmentation of the interface generated from the openapi spec.
func New(client *itunes.Client, logger commonlogger.ErrorInfoDebugger) http.Handler {
	r := chi.NewMux()

	r.Route("/v1", func(r chi.Router) {
		// init request/response middleware
		r.Use(logging.LoggingMiddleware(logger))

		HandlerFromMux(handlers{Client: client, logger: logger}, r)
		r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK")) //nolint: errcheck #nothing be be gained from this check
		})
	})
	return r
}

type handlers struct {
	Client *itunes.Client
	logger commonlogger.ErrorInfoDebugger
}

// GetArtistArtistAlbumTitle serves as the artist/:artist/album/:title endpoint
func (s handlers) GetArtistArtistAlbumTitle(w http.ResponseWriter, r *http.Request, artist string, title string, params GetArtistArtistAlbumTitleParams) {
	w.Header().Add("Content-Type", "application/json")

	artist = normalise(artist)
	title = normalise(title)
	// search itunes
	searchTerm := fmt.Sprintf("%s - %s", artist, title)
	out, err := s.Client.Search(r.Context(), searchTerm, "gb", "album")

	if err != nil {
		s.writeErrorResponse(r.Context(), w, WrapError(err, ErrBadGateway))
		return
	}

	size := fetchQueryParameter(r, "size")

	rtn := []domain.Album{}

	for _, item := range out.Results {
		artistMatch := normalise(item.ArtistName) == artist

		if !artistMatch || !strings.Contains(normalise(item.CollectionName), title) {
			s.logger.Debug("dropping result", item)
			continue
		}

		row := domain.Album{
			ImageURL:   item.ArtworkURL100,
			Title:      item.CollectionName,
			ArtistName: item.ArtistName,
		}

		if len(size) > 0 {
			// Update return itunes return values to match the size requested by OUR client
			row.ImageURL = strings.Replace(row.ImageURL, "100x100", fmt.Sprintf("%sx%s", size[0], size[0]), 1)
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

// fetchQueryParameter grabs parameter from the url. The query parameters, not the route parameters
func fetchQueryParameter(r *http.Request, key string) []string {
	if v, ok := r.URL.Query()[key]; ok {
		return v
	}
	return nil
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
