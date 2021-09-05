package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/env"

	"github.com/johnmackenzie91/commonlogger"
)

type HTTP struct {
	server http.Server
	logger commonlogger.ErrorInfoDebugger
}

// New creates a new server from environment config
func New(e env.Config, handler http.Handler, logger commonlogger.ErrorInfoDebugger) HTTP {
	return HTTP{
		server: http.Server{
			Addr:              e.HTTPAddress,
			Handler:           handler,
			ReadTimeout:       time.Duration(e.HTTPReadTimeout) * time.Second,
			ReadHeaderTimeout: time.Duration(e.HTTPReadTimeout) * time.Second,
			WriteTimeout:      time.Duration(e.HTTPWriteTimeout) * time.Second,
		},
		logger: logger,
	}
}

// ListenAndServe wraps listening of the server functionality.
// If an error occurs during execution the error is returned on the out channel.
// If server closes gracefully, the out channel is closed.
func (h *HTTP) ListenAndServe() chan error {
	h.logger.Info("running server")
	out := make(chan error, 1)
	go func() {
		if err := h.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			out <- fmt.Errorf("failed to close server: %w", err)
		}
		close(out)
	}()
	return out
}

// Shutdown shuts the server down
func (h *HTTP) Shutdown(ctx context.Context) error {
	h.logger.Info("shutting down server")
	return h.server.Shutdown(ctx)
}
