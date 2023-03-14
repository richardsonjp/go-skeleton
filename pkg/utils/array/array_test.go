package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteEmpty(t *testing.T) {
	testCases := []struct{ in, want []string }{
		{[]string{"", "a", "b", "c", "", "d"}, []string{"a", "b", "c", "d"}},
		{[]string{"", "a", "c", "", "d"}, []string{"a", "c", "d"}},
		{[]string{"", "a", "", "", ""}, []string{"a"}},
		{[]string{"", "", ""}, []string(nil)},
	}

	for _, tc := range testCases {
		result := DeleteEmpty(tc.in)
		assert.Equal(t, tc.want, result, "DeleteEmpty: should %v be equal to %v", result, tc.want)
	}
}

func TestContainsUint(t *testing.T) {
	testCases := []struct {
		in     []uint
		search uint
		want   bool
	}{
		{[]uint{1, 2, 3, 4, 5}, 3, true},
		{[]uint{1, 2, 3, 4, 5}, 10, false},
	}

	for _, tc := range testCases {
		result := ContainsUint(tc.in, tc.search)
		assert.Equal(t, tc.want, result, "ContainsUint: search %v in %v should return %v", tc.search, tc.in, tc.want)
	}
}
