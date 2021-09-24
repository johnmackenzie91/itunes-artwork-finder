package itunes

import (
	"context"
	"net/http"
	"testing"

	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/pkg/itunes/mocks"

	"github.com/johnmackenzie91/httptestfixtures"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//go:generate mockery -name Doer
func TestClient_Search(t *testing.T) {
	stubCtx := context.Background()

	tests := []struct {
		name        string
		setUpMocks  func(client *mocks.Doer)
		FixtureFile string
		expected    SearchResponse
		error       error
	}{
		{
			name:        "results returned",
			FixtureFile: "./testdata/sams_town",
			expected: SearchResponse{
				ResultCount: 2,

				Results: []Result{
					{
						WrapperType:    "collection",
						CollectionType: "Album",
						CollectionName: "Sam's Town",
						ArtistName:     "The Killers",
						ArtworkURL100:  "https://is4-ssl.mzstatic.com/image/thumb/Music124/v4/ed/76/b3/ed76b36d-cbf7-3518-b243-c2fdce84f53f/source/100x100bb.jpg",
						ArtworkURL60:   "https://is4-ssl.mzstatic.com/image/thumb/Music124/v4/ed/76/b3/ed76b36d-cbf7-3518-b243-c2fdce84f53f/source/60x60bb.jpg",
					},
					{
						WrapperType:    "collection",
						CollectionType: "Album",
						CollectionName: "Sam's Town",
						ArtistName:     "The Killers",
						ArtworkURL100:  "https://is2-ssl.mzstatic.com/image/thumb/Music125/v4/3f/29/f0/3f29f099-2e2e-7c26-52f6-600f1ad405c3/source/100x100bb.jpg",
						ArtworkURL60:   "https://is2-ssl.mzstatic.com/image/thumb/Music125/v4/3f/29/f0/3f29f099-2e2e-7c26-52f6-600f1ad405c3/source/60x60bb.jpg",
					},
				},
			},
		},
		{
			name:        "no results returned",
			FixtureFile: "./testdata/no_results",
			error:       nil,
			expected: SearchResponse{
				ResultCount: 0,
				Results:     []Result{},
			},
		},
		{
			name:        "bad response",
			FixtureFile: "./testdata/bad_request",
			error:       errNon2XX(http.StatusBadRequest),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := mocks.Doer{}

			res := httptestfixtures.MustLoadRequest(t, tt.FixtureFile)
			mockClient.On("Do", mock.Anything).Return(res, nil)

			c := Client{
				client: &mockClient,
				logger: logrus.New(),
			}
			got, err := c.Search(stubCtx, "n/a", "n/a", "n/a")
			assert.Equal(t, tt.expected, got)
			assert.ErrorIs(t, err, tt.error)
			assert.Nil(t, res.Body.Close())
		})
	}
}
