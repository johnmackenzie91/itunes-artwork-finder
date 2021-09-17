package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRestV1AlbumSearchParams_Validate(t *testing.T) {
	testCases := []struct {
		desc     string
		input    GetRestV1AlbumSearchParams
		expected error
	}{
		{
			desc:     "1. required title parameter missing",
			input:    GetRestV1AlbumSearchParams{},
			expected: errMissingParameters([]string{"title"}),
		},
		{
			desc: "2. no error for missing non required parameters",
			input: GetRestV1AlbumSearchParams{
				Title: "some title",
			},
			expected: nil,
		},
		{
			desc: "3. all parameters provided",
			input: GetRestV1AlbumSearchParams{
				Artist: func(input string) *string { return &input }("some artist"),
				Title:  "some title",
				Size:   func(input int) *int { return &input }(400),
			},
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.input.Validate()
			assert.Equal(t, tc.expected, err)
		})
	}
}
