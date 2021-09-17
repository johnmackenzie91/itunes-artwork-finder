package app

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/app/mocks"
	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/finder"

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
			expectedStatus: http.StatusBadGateway,
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

var (
	stubRequest, _ = http.NewRequest("GET", "http://stub.stub", nil)
	stubCtx        = context.Background()
	errTest        = errors.New("test error")
	emptyResponse  = finder.SearchResponse{}
)

func Test_handlers_GetRestV1AlbumSearch(t *testing.T) {
	testCases := []struct {
		desc            string
		inputParameters GetRestV1AlbumSearchParams
		setupMocks      func(mock *mocks.Searcher)
		expectedOut     string
		expectedCode    int
	}{
		{
			desc: "1. happy path, results returned",
			inputParameters: GetRestV1AlbumSearchParams{
				Title: "some title",
			},
			setupMocks: func(mock *mocks.Searcher) {
				res := finder.SearchResponse{
					ResultCount: 1,
					Results: []finder.Result{
						{
							Artist:         "some artist",
							CollectionName: "some title",
							Link:           "http://some-artist.jpeg",
						},
					},
				}
				mock.On("Search", stubCtx, "some title", "gb", "album").Return(res, nil)
			},
			expectedCode: http.StatusOK,
			expectedOut:  `[{"title":"some title","artist_name":"some artist","image_url":"http://some-artist.jpeg"}]`,
		},
		{
			desc: "2. error returned from fetcher",
			inputParameters: GetRestV1AlbumSearchParams{
				Title: "some title",
			},
			setupMocks: func(mock *mocks.Searcher) {
				mock.On("Search", stubCtx, "some title", "gb", "album").Return(emptyResponse, errTest)
			},
			expectedCode: http.StatusBadGateway,
			expectedOut:  `{"msg":"bad gateway"}`,
		},
		{
			desc:            "3. required artist parameter missing",
			inputParameters: GetRestV1AlbumSearchParams{},
			setupMocks:      func(mock *mocks.Searcher) {},
			expectedCode:    http.StatusBadRequest,
			expectedOut:     `{"msg":"bad request"}`,
		},
		{
			desc: "4. title is provided and results that do not match title are removed",
			inputParameters: GetRestV1AlbumSearchParams{
				Artist: func(input string) *string { return &input }("some artist"),
				Title:  "album two",
			},
			setupMocks: func(mock *mocks.Searcher) {
				res := finder.SearchResponse{
					ResultCount: 1,
					Results: []finder.Result{
						{
							Artist:         "some artist",
							CollectionName: "album one",
							Link:           "http://some-artist.jpeg",
						},
						{
							Artist:         "some artist",
							CollectionName: "album two",
							Link:           "http://album-two.jpeg",
						},
					},
				}
				mock.On("Search", stubCtx, "some artist - album two", "gb", "album").Return(res, nil)
			},
			expectedCode: http.StatusOK,
			expectedOut:  `[{"title":"album two","artist_name":"some artist","image_url":"http://album-two.jpeg"}]`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// arrange
			searcher := mocks.Searcher{}
			tc.setupMocks(&searcher)

			// act
			sut := handlers{
				Client: &searcher,
				logger: logrus.New(),
			}
			w := httptest.NewRecorder()
			sut.GetRestV1AlbumSearch(w, stubRequest, tc.inputParameters)

			// assert
			assert.Equal(t, tc.expectedCode, w.Code)
			assert.Equal(t, tc.expectedOut, w.Body.String())
			searcher.AssertExpectations(t)
		})
	}
}
