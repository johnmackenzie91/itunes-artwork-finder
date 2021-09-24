package main

import (
	"net/http"
	"os"
	"os/signal"
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

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

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

	// wait until server shut down or os interrupts
	select {
	case <-quit:
		log.Info("OS Interrupt ....")
		log.Info("closing down server ...")
	case err, open := <-svr.ListenAndServe():
		if open {
			log.Error(err)
			break
		}
		log.Info("closing down server...")
	}
}
