package app

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_filterResults(t *testing.T) {

}

func Test_handlers_writeErrorResponse(t *testing.T) {
	// arrange
	stubCtx := context.Background()
	errTest := errors.New("test error")

	testCases := []struct {
		name           string
		input          error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "bad request sentinel err",
			input:          WrapError(errTest, ErrBadGateway),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "{\"msg\":\"bad gateway\"}",
		},
		{
			name:           "an unwrapped sentinel error",
			input:          ErrNotFound,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "{\"msg\":\"not found\"}",
		},
		{
			name:           "standard go error",
			input:          errTest,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "{\"msg\":\"internal server error\"}",
		},
	}

	// act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := handlers{
				logger: logrus.New(),
			}

			rec := httptest.NewRecorder()
			sut.writeErrorResponse(stubCtx, rec, tc.input)

			// assert
			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.Equal(t, tc.expectedBody, rec.Body.String())
		})
	}
}
