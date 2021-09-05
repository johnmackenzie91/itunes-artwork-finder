package itunes

import (
	"context"
	"net/http"
	"testing"

	"bitbucket.org/johnmackenzie91/itunes-artwork-proxy-api/internal/itunes/mocks"

	"github.com/johnmackenzie91/httptestfixtures"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//go:generate mockery -name Doer
//go:generate curl -i https://itunes.apple.com/search?country=gb&entity=album&term=the+killers+-+sams+town -o ./testdata/sams_town
//go:generate curl -i https://itunes.apple.com/search?country=gb\u0026entity=album\u0026term=the+killers+-+sams+town -o ./testdata/bad_request
func TestClient_Search(t *testing.T) {
	stubCtx := context.Background()

	tests := []struct {
		name       string
		setUpMocks func(client *mocks.Doer)
		expected   SearchResponse
		error      error
	}{
		{
			name: "results returned",
			setUpMocks: func(client *mocks.Doer) {
				res := httptestfixtures.MustLoadRequest(t, "./testdata/sams_town")
				client.On("Do", mock.Anything).Return(res, nil)
			},
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
			name: "bad response",
			setUpMocks: func(client *mocks.Doer) {
				res := httptestfixtures.MustLoadRequest(t, "./testdata/bad_request")
				client.On("Do", mock.Anything).Return(res, nil)
			},
			error: errNon2XX(http.StatusBadRequest),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := mocks.Doer{}
			tt.setUpMocks(&mockClient)

			c := Client{
				client: &mockClient,
				logger: logrus.New(),
			}
			got, err := c.Search(stubCtx, "n/a", "n/a", "n/a")
			assert.Equal(t, tt.expected, got)
			assert.ErrorIs(t, err, tt.error)
		})
	}
}
