package redoc

import (
	"net/http"
)

// V1Docs returns the page that will host the specifiction
func V1Docs() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

// V1Spec returns the sepcification for v1 version of the api
func V1Spec() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
