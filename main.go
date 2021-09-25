package main

import (
	"net/http"
	"time"

	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/app"
	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/artwork"
	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/env"
	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/logger"
	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/server"
)

func main() {
	// Load environment variables
	e := env.MustLoad()

	log := logger.New(e)

	// init clients that speak to downstream services
	finderAdapter, err := artwork.Itunes(
		e.ItunesEndpoint,
		http.Client{Timeout: 5 * time.Second},
		log,
	)

	if err != nil {
		panic(err)
	}

	// Init router and app
	a := app.New(finderAdapter, log)
	svr := server.New(e, a, log)

	// ListenAndServe hangs until server is closed, or error received
	if err := svr.ListenAndServe(); err != nil {
		panic(err)
	}

	if err := svr.Shutdown(); err != nil {
		panic(err)
	}
}
