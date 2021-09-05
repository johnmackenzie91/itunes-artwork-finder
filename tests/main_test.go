package tests

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
)

var long = flag.Bool("long", false, "if present run the long, integration tests")

func TestMain(m *testing.M) {
	if long == nil {
		return
	}

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	res, err := pool.RunWithOptions(sutContainer.options)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err := pool.Retry(sutContainer.HealthCheck); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	code := m.Run()

	if err := pool.Purge(res); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

var httpClient = &http.Client{}

func Test_Search_Endpoint(t *testing.T) {
	testCases := []struct {
		name           string
		request        *http.Request
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Get request with real info",
			request: func() *http.Request {
				url := sutContainer.URL("arctic monkeys", "humbug")
				r, _ := http.NewRequest("GET", url, strings.NewReader(""))
				return r
			}(),
			expectedStatus: http.StatusOK,
			// load fixture from file
			expectedBody: func() string {
				got, err := ioutil.ReadFile("./testdata/search_get_humbug.json")
				if err != nil {
					t.Fatal(err)
				}
				return string(got)
			}(),
		},
		{
			name: "not found",
			request: func() *http.Request {
				url := sutContainer.URL("no-artist", "unknown")
				r, _ := http.NewRequest("GET", url, strings.NewReader(""))
				return r
			}(),
			expectedStatus: http.StatusNotFound,
			// load fixture from file
			expectedBody: "{\"msg\":\"not found\"}",
		},
		{
			name: "POST request",
			request: func() *http.Request {
				url := sutContainer.URL("nothing", "nothing")
				r, _ := http.NewRequest("POST", url, strings.NewReader(""))
				return r
			}(),
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// attempt request
			res, err := httpClient.Do(tc.request)
			if err != nil {
				t.Fatal(err)
			}

			// assert on response
			assert.Equal(t, tc.expectedStatus, res.StatusCode)
			b, err := ioutil.ReadAll(res.Body)
			assert.Nil(t, err)
			assert.Equal(t, tc.expectedBody, string(b))

			res.Body.Close() //nolint: errcheck #nothing be be gained from this check
		})
	}
}
