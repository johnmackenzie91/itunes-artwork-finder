package itunes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetDomain(t *testing.T) {
	testCases := []struct {
		desc        string
		input       string
		expected    string
		expectedErr error
	}{
		{
			desc:        "1. blank input defaults to prod url",
			input:       "",
			expected:    "https://itunes.apple.com/search",
			expectedErr: nil,
		},
		{
			desc:        "2. bad string returns error",
			input:       "  ",
			expectedErr: errBadURL("  "),
		},
		{
			desc:        "3. bad url received",
			input:       "bad UrL",
			expectedErr: errBadURL("bad UrL"),
		},
	}
	for _, tc := range testCases {
		// arrange
		sut := SetDomain(tc.input)
		input := Client{}
		// act
		err := sut(&input)
		// assert
		assert.Equal(t, tc.expectedErr, err)
		assert.Equal(t, tc.expected, input.domain)
	}
}
