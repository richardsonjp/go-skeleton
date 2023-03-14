package strings

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func stringToPointer(text string) *string {
	return &text
}

func TestMaskString(t *testing.T) {
	testCases := []struct {
		in struct {
			s         string
			startFrom int
		}
		want string
	}{
		{
			in: struct {
				s         string
				startFrom int
			}{"hello", 1},
			want: "h****",
		},
		{
			in: struct {
				s         string
				startFrom int
			}{"hello world", 1},
			want: "h**********",
		},
		{
			in: struct {
				s         string
				startFrom int
			}{"h", 1},
			want: "h",
		},
		{
			in: struct {
				s         string
				startFrom int
			}{"", 1},
			want: NULL,
		},
		{
			in: struct {
				s         string
				startFrom int
			}{"hello", -1},
			want: NULL,
		},
	}

	for i, tc := range testCases {
		result := MaskString(tc.in.s, tc.in.startFrom)

		assert.Equal(t, tc.want, result, fmt.Sprintf("MaskString: should be equals at index: %v", i))
	}
}

func TestMaskIdentityCardNumber(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{
			in:   "1234567890123456",
			want: "12345678901*****",
		},
		{
			in:   "hello world",
			want: "hello *****",
		},
		{
			in:   "hello",
			want: NULL,
		},
		{
			in:   "worl",
			want: NULL,
		},
	}

	for i, tc := range testCases {
		result := MaskIdentityCardNumber(tc.in)

		assert.Equal(t, tc.want, result, fmt.Sprintf("MaskIdentityCardNumber: should be equals at index: %v", i))
	}
}

func TestMaskUUIDV4(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{
			in:   "1234567890123456",
			want: "12345678********",
		},
		{
			in:   "hello world",
			want: "hello wo***",
		},
		{
			in:   "hello",
			want: NULL,
		},
		{
			in:   "worl",
			want: NULL,
		},
	}

	for i, tc := range testCases {
		result := MaskUUIDV4(tc.in)

		assert.Equal(t, tc.want, result, fmt.Sprintf("MaskUUIDV4: should be equals at index: %v", i))
	}
}

func TestConvertMapToString(t *testing.T) {
	testCases := []struct {
		in   map[string]string
		want string
	}{
		{
			in:   map[string]string{"asdasd": "qweqwe"},
			want: "map[asdasd:qweqwe]",
		},
		{
			in:   map[string]string{},
			want: "",
		},
	}

	for i, tc := range testCases {
		result := ConvertMapToString(tc.in)

		assert.Equal(t, tc.want, result, fmt.Sprintf("ConvertMapToString: should be equals at index: %v", i))
	}
}

func TestDeleteEmptyElement(t *testing.T) {
	testCases := []struct {
		in   []string
		want []string
	}{
		{
			in:   []string{"a", "b", "c"},
			want: []string{"a", "b", "c"},
		},
		{
			in:   []string{"a", " ", "b", "c"},
			want: []string{"a", " ", "b", "c"},
		},
		{
			in:   []string{"a", "", "b", "", "c"},
			want: []string{"a", "b", "c"},
		},
	}

	for i, tc := range testCases {
		result := DeleteEmptyElement(tc.in)

		assert.Equal(t, tc.want, result, fmt.Sprintf("DeleteEmptyElement: should be equals at index: %v", i))
	}
}

func TestBytesToBase64String(t *testing.T) {
	testCases := []struct {
		in   []byte
		want string
	}{
		{
			in:   []byte("hello"),
			want: "aGVsbG8=",
		},
		{
			in:   []byte("hello world"),
			want: "aGVsbG8gd29ybGQ=",
		},
	}

	for i, tc := range testCases {
		result := BytesToBase64String(tc.in)

		assert.Equal(t, tc.want, result, fmt.Sprintf("BytesToBase64String: should be equals at index: %v", i))
	}
}

func TestReplaceFormatIndex(t *testing.T) {
	testCases := []struct {
		in struct {
			textFormat  *string
			substitutes []string
		}
		want *string
	}{
		{
			in: struct {
				textFormat  *string
				substitutes []string
			}{textFormat: stringToPointer("The {0} brown fox {1} the desert"), substitutes: []string{"quick", "at"}},
			want: stringToPointer("The quick brown fox at the desert"),
		},
		{

			in: struct {
				textFormat  *string
				substitutes []string
			}{textFormat: nil, substitutes: nil},
			want: nil,
		},
		{

			in: struct {
				textFormat  *string
				substitutes []string
			}{textFormat: stringToPointer("The {0} brown fox {1} the desert"), substitutes: nil},
			want: stringToPointer("The {0} brown fox {1} the desert"),
		},
		{

			in: struct {
				textFormat  *string
				substitutes []string
			}{textFormat: stringToPointer("The {0} brown fox {1} the desert"), substitutes: []string{}},
			want: stringToPointer("The {0} brown fox {1} the desert"),
		},
	}

	for i, tc := range testCases {
		result := ReplaceFormatIndex(tc.in.textFormat, tc.in.substitutes...)

		assert.Equal(t, tc.want, result, fmt.Sprintf("ReplaceFormatIndex: should be equals at index: %v", i))
	}
}

func TestFirstUpperCase(t *testing.T) {
	testCases := []struct {
		in   *string
		want *string
	}{
		{
			in:   stringToPointer("1234567890123456"),
			want: stringToPointer("1234567890123456"),
		},
		{
			in:   stringToPointer("hello world"),
			want: stringToPointer("Hello World"),
		},
		{
			in:   stringToPointer("HELLO WORLD"),
			want: stringToPointer("Hello World"),
		},
		{
			in:   nil,
			want: nil,
		},
	}

	for i, tc := range testCases {
		result := FirstUpperCase(tc.in)

		assert.Equal(t, tc.want, result, fmt.Sprintf("FirstUpperCase: should be equals at index: %v", i))
	}
}

func TestLowerCase(t *testing.T) {
	testCases := []struct {
		in   *string
		want *string
	}{
		{
			in:   stringToPointer("1234567890123456"),
			want: stringToPointer("1234567890123456"),
		},
		{
			in:   stringToPointer("hello world"),
			want: stringToPointer("hello world"),
		},
		{
			in:   stringToPointer("Hello World"),
			want: stringToPointer("hello world"),
		},
		{
			in:   stringToPointer("HELLO WORLD"),
			want: stringToPointer("hello world"),
		},
		{
			in:   nil,
			want: nil,
		},
	}

	for i, tc := range testCases {
		result := LowerCase(tc.in)

		assert.Equal(t, tc.want, result, fmt.Sprintf("LowerCase: should be equals at index: %v", i))
	}
}

func TestUnifyLatLong(t *testing.T) {
	testCases := []struct {
		in struct {
			latitude  *string
			longitude *string
		}
		want *string
	}{
		{
			in: struct {
				latitude  *string
				longitude *string
			}{latitude: stringToPointer("1234567890123456 "), longitude: stringToPointer("1234567890123456 ")},
			want: stringToPointer("1234567890123456, 1234567890123456"),
		},
		{
			in: struct {
				latitude  *string
				longitude *string
			}{latitude: stringToPointer(" "), longitude: stringToPointer(" ")},
			want: nil,
		},
		{
			in: struct {
				latitude  *string
				longitude *string
			}{latitude: stringToPointer(""), longitude: stringToPointer("")},
			want: nil,
		},
		{
			in: struct {
				latitude  *string
				longitude *string
			}{latitude: stringToPointer("123"), longitude: stringToPointer("")},
			want: nil,
		},
		{
			in: struct {
				latitude  *string
				longitude *string
			}{latitude: stringToPointer(""), longitude: stringToPointer("123")},
			want: nil,
		},
		{
			in: struct {
				latitude  *string
				longitude *string
			}{latitude: nil, longitude: stringToPointer("123")},
			want: nil,
		},
		{
			in: struct {
				latitude  *string
				longitude *string
			}{latitude: stringToPointer("123"), longitude: nil},
			want: nil,
		},
		{
			in: struct {
				latitude  *string
				longitude *string
			}{latitude: nil, longitude: nil},
			want: nil,
		},
	}

	for i, tc := range testCases {
		result := UnifyLatLong(tc.in.latitude, tc.in.longitude)

		assert.Equal(t, tc.want, result, fmt.Sprintf("UnifyLatLong: should be equals at index: %v", i))
	}
}
