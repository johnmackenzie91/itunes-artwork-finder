package main

import (
	"os"
	"os/signal"

	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/app"
	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/env"
	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/logger"
	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/server"
	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/pkg/itunes"
)

func main() {
	// Load environment variables
	e := env.MustLoad()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	log := logger.New(e)

	// init clients that speak to downstream services
	itunesCli := itunes.New(log, e)

	// Init router and app
	a := app.New(itunesCli, log)
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
