package logging

import (
	"net/http"

	"github.com/johnmackenzie91/commonlogger"
)

// LoggingMiddleware logs the incoming HTTP request & its duration.
func LoggingMiddleware(logger commonlogger.ErrorInfoDebugger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					logger.Error(r.Context(), err)
				}
			}()

			// log the request
			logger.Info(r.Context(), "request received", r)

			wrapped := wrapResponseWriter(w)

			// run handler
			next.ServeHTTP(wrapped, r)

			// log response
			logger.Info(r.Context(), wrapped, "response sent")
		})

	}
}
