package wording

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func float64ToPointer(value float64) *float64 {
	return &value
}

func TestFormatIDRCurrency(t *testing.T) {
	testCases := []struct {
		in   *float64
		want string
	}{
		{in: float64ToPointer(500000), want: "Rp500.000"},
		{in: float64ToPointer(50000), want: "Rp50.000"},
		{in: float64ToPointer(5000), want: "Rp5.000"},
		{in: float64ToPointer(500), want: "Rp500"},
	}

	for index, tc := range testCases {
		result := FormatIDRCurrency(tc.in)
		assert.Equal(t, tc.want, result, "FormatIDRCurrency: should be equals at index %v", index)
	}
}

func TestDenormalizePhoneNumber(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{in: "+628123456789  ", want: "08123456789"},
		{in: "+08123456789", want: "08123456789"},
		{in: "628123456789", want: "08123456789"},
		{in: "6281234567", want: "081234567"},
		{in: "6281234567", want: "081234567"},
		{in: "+", want: ""},
		{in: "+62", want: "0"},
		{in: "62", want: "62"},
	}

	for index, tc := range testCases {
		result := DenormalizePhoneNumber(tc.in)
		assert.Equal(t, tc.want, result, "DenormalizePhoneNumber: should be equals at index %v", index)
	}
}

func TestNormalizePhoneNumber(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{in: "08123456789  ", want: "628123456789"},
		{in: "08123456789", want: "628123456789"},
		{in: "62123456789", want: "62123456789"},
		{in: "1", want: "1"},
	}

	for index, tc := range testCases {
		result := NormalizePhoneNumber(tc.in)
		assert.Equal(t, tc.want, result, "NormalizePhoneNumber: should be equals at index %v", index)
	}
}

func TestToSnakeCase(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{in: "helloWorld", want: "hello_world"},
		{in: "helloWorld123", want: "hello_world123"},
		{in: "helloWorldOne", want: "hello_world_one"},
		{in: "helloWorld.One", want: "hello_world._one"},
		{in: "helloWorld.one", want: "hello_world.one"},
		{in: "hello123", want: "hello123"},
		{in: "123hello", want: "123hello"},
		{in: "hello world", want: "hello world"},
	}

	for index, tc := range testCases {
		result := ToSnakeCase(tc.in)
		assert.Equal(t, tc.want, result, "ToSnakeCase: should be equals at index %v", index)
	}
}

func TestFormatWords(t *testing.T) {
	testCases := []struct {
		in   float64
		want string
	}{
		{in: -1, want: "negatif satu"},
		{in: 1, want: "satu"},
		{in: 10, want: "sepuluh"},
		{in: 100, want: "seratus"},
		{in: 1000, want: "seribu"},
		{in: 10000, want: "sepuluh ribu"},
		{in: 100000, want: "seratus ribu"},
		{in: 1000000, want: "satu juta"},
		{in: 10000000, want: "sepuluh juta"},
		{in: 11111111, want: "sebelas juta seratus sebelas ribu seratus sebelas"},
		{in: 11111109, want: "sebelas juta seratus sebelas ribu seratus sembilan"},
		{in: 1342345567, want: "satu miliar tiga ratus empat puluh dua juta tiga ratus empat puluh lima ribu lima ratus enam puluh tujuh"},
		{in: 100000000000, want: "seratus miliar"},
		{in: 1000000000000, want: "satu triliun"},
		{in: 10000000000001, want: "sepuluh triliun satu"},
		{in: 1000000000000111, want: "satu kuadtriliun seratus sebelas"},
	}

	for index, tc := range testCases {
		result := FormatWords(tc.in)
		assert.Equal(t, tc.want, result, "FormatWords: should be equals at index %v", index)
	}
}
