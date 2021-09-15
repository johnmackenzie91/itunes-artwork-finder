package redoc

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

// TestHandlers_V1Docs I expect the contents of ./docs/redoc.html to be returned from handler
func TestHandlers_V1Docs(t *testing.T) {
	// arrange
	h := Handlers{
		logger: logrus.New(),
	}
	w := httptest.NewRecorder()

	// act
	h.V1Docs(w, nil)

	// assert
	expected, err := os.ReadFile("./docs/redoc.html")
	assert.Nil(t, err)
	assert.Equal(t, string(expected), w.Body.String())
}

// TestHandlers_V1Spec I expect the contents of ./docs/openapi.json to be returned from handler
func TestHandlers_V1Spec(t *testing.T) {
	// arrange
	h := Handlers{
		logger: logrus.New(),
	}
	w := httptest.NewRecorder()

	// act
	h.V1Spec(w, nil)

	// assert
	expected, err := os.ReadFile("./docs/openapi.json")
	assert.Nil(t, err)
	assert.Equal(t, string(expected), w.Body.String())
}
