package logger

import (
	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/env"
	"github.com/johnmackenzie91/commonlogger"
	"github.com/sirupsen/logrus"
	"github.com/johnmackenzie91/commonlogger/resolvers"
)

func New(env env.Config) commonlogger.ErrorInfoDebugger {
	l := logrus.New()

	lvl, err := logrus.ParseLevel(env.LogLevel)
	if err != nil {
		lvl = logrus.InfoLevel
	}
	l.SetLevel(lvl)
	l.SetFormatter(&logrus.JSONFormatter{})
	return commonlogger.New(l, commonlogger.Config{
		RequestResolver:  resolvers.ResolveJSONRequest,
		ResponseResolver: resolvers.ResolveJSONResponse,
	})
}