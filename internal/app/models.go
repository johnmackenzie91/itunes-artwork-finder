package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var (
	ErrNotFound   = sentinelAPIError{Code: http.StatusNotFound, Msg: "not found"}
	ErrInternal   = sentinelAPIError{Code: http.StatusBadRequest, Msg: "internal server error"}
	ErrBadGateway = sentinelAPIError{Code: http.StatusBadGateway, Msg: "bad gateway"}
	ErrBadRequest = sentinelAPIError{Code: http.StatusBadRequest, Msg: "bad request"}
)

type APIError interface {
	APIError() (int, string)
}

// APIError i want to limit the errors returned to the user.
// Not to give away too much information on what has gone wrong.
type sentinelAPIError struct {
	Msg  string `json:"msg"`
	Code int    `json:"-"`
}

// Error fulfills the error interface
func (e sentinelAPIError) Error() string {
	return e.Msg
}

func (e sentinelAPIError) APIError() (int, string) {
	return e.Code, e.Msg
}

// JSON returns the error in json format
func (e sentinelAPIError) JSON() []byte {
	// json.Marshal only fails when we attempt to marshal channels and function values.
	// along with cyclic data structures. We can assume this will not error because we own the data type
	b, err := json.Marshal(e)

	if err != nil {
		panic(err)
	}
	return b
}

// sentinelWrappedError helps associate a sentinel error to an internal error.
type sentinelWrappedError struct {
	error
	sentinel sentinelAPIError
}

func (e sentinelWrappedError) Is(err error) bool {
	return e.sentinel == err
}

func (e sentinelWrappedError) APIError() (int, string) {
	return e.sentinel.APIError()
}

// WrapError wraps an internal possibly sensitive error into a sentinel error
func WrapError(err error, sentinel sentinelAPIError) error {
	return sentinelWrappedError{error: err, sentinel: sentinel}
}

type errMissingParameters []string

func (e errMissingParameters) Error() string {
	return fmt.Sprintf("required parameters missing: %s", strings.Join(e, ", "))
}

// Validate ensures that all the correct properties are set
func (p GetRestV1AlbumSearchParams) Validate() error {
	missing := []string{}
	if p.Title == "" {
		missing = append(missing, "title")
		return errMissingParameters(missing)
	}
	return nil
}
