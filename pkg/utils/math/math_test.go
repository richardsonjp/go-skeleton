package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertStringToUint(t *testing.T) {
	testCases := []struct {
		in   string
		want uint
	}{
		{"a", 0},
		{"1", 1},
		{"11", 11},
		{"32", 32},
		{"64", 64},
		{"11111111111111111111111111", 0},
	}

	for index, tc := range testCases {
		result, err := ConvertStringToUint(tc.in)

		if err == nil {
			assert.Equal(t, tc.want, result, "ConvertStringToUint: index at %v", index)
			assert.ErrorIs(t, nil, err)
		} else {
			assert.Equal(t, tc.want, result, "ConvertStringToUint: index at %v", index)
			assert.Error(t, err)
		}
	}
}
