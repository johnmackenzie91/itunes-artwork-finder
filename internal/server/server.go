package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
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
func (h *HTTP) ListenAndServe() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	h.logger.Info("running server")
	out := make(chan error, 1)
	go func() {
		if err := h.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			out <- fmt.Errorf("failed to close server: %w", err)
		}
		close(out)
	}()

	var err error
	// wait until server shut down or os interrupts
	select {
	case <-quit:
		h.logger.Info("OS Interrupt ....")
		h.logger.Info("closing down server ...")
	case err, open := <-out:
		if open {
			h.logger.Error(err)
			break
		}
		h.logger.Info("closing down server...")
	}

	return err
}

// Shutdown shuts the server down
func (h *HTTP) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	h.logger.Info("shutting down server")
	return h.server.Shutdown(ctx)
}
