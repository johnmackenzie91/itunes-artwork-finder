package itunes

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var stubCtx = context.Background()

func TestClient_buildSearchRequest(t *testing.T) {
	// arrange
	sut := Client{}

	// act
	out, err := sut.buildSearchRequest(stubCtx, "some artist - some album", "GB", "album")

	// assert
	assert.Nil(t, err)
	assert.Equal(t, "?country=GB&entity=album&term=some+artist+-+some+album", out.URL.String())
}
