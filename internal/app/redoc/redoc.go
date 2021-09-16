package redoc

import (
	"net/http"

	"embed"
	"github.com/johnmackenzie91/commonlogger"
)

//go:embed docs/*
var docs embed.FS

type Handlers struct {
	logger commonlogger.ErrorInfoDebugger
}

func New(logger commonlogger.ErrorInfoDebugger) Handlers {
	return Handlers{

		logger: logger,
	}
}

// V1Docs returns the page that will host the specifiction
func (h Handlers) V1Docs(w http.ResponseWriter, r *http.Request) {
	f, err := docs.ReadFile("docs/redoc.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Error(err)
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(f); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Error(err)
	}
}

// V1Spec returns the sepcification for v1 version of the api
func (h Handlers) V1Spec(w http.ResponseWriter, r *http.Request) {
	f, err := docs.ReadFile("docs/openapi.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Error(err)
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(f); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Error(err)
	}
}
