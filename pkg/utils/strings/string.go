package strings

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"go-skeleton/pkg/utils/null"
	"strings"
)

const NULL = "NULL"

func MaskString(s string, startFrom int) string {
	// if startFrom is negative or the "subject" have less length than startFrom, will return NULL
	if startFrom <= 0 || len(s) < startFrom {
		return NULL
	}

	rs := []rune(s)
	for i := startFrom; i < len(rs); i++ {
		rs[i] = '*'
	}
	return string(rs)
}

func MaskRight(s string, minimum int) string {
	startFrom := len(s) - minimum
	return MaskString(s, startFrom)
}

// Function to mask identity card number
func MaskIdentityCardNumber(s string) string {
	maskLength := 5
	start := len(s) - maskLength
	return MaskString(s, start)
}

func MaskUUIDV4(s string) string {
	return MaskString(s, 8)
}

func ConvertMapToString(value map[string]string) string {
	if len(value) < 1 {
		return ""
	}
	str := "map["
	for key, val := range value {
		str = fmt.Sprintf("%s%s:%s", str, key, val)
	}
	return fmt.Sprintf("%s]", str)
}

func DeleteEmptyElement(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

// Converting byte array into base 64 string
func BytesToBase64String(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// Replace string format with subtitutes.
// @param textFormat example "{0} bla bla {1} bli bli {2} blu blu {3}"
// @param subtitutes
// @Return string | null
func ReplaceFormatIndex(textFormat *string, subtitutes ...string) *string {
	if textFormat == nil {
		return textFormat
	}
	if subtitutes == nil {
		return textFormat
	}
	if len(subtitutes) == 0 {
		return textFormat
	}
	var newStringBuffer bytes.Buffer
	newStringBuffer.WriteString(*textFormat)
	arrayText := subtitutes
	for i, v := range arrayText {
		newText := strings.Replace(newStringBuffer.String(), "{"+fmt.Sprint(i)+"}", v, -1)
		newStringBuffer.Reset()
		newStringBuffer.WriteString(newText)
	}
	newString := newStringBuffer.String()
	return &newString
}

func FirstUpperCase(s *string) *string {
	if null.IsNil(s) {
		return s
	}
	value := strings.Title(strings.ToLower(*s))
	return &value
}

func LowerCase(s *string) *string {
	if null.IsNil(s) {
		return s
	}
	value := strings.ToLower(*s)
	return &value
}

func UnifyLatLong(latitude *string, longitude *string) *string {
	if null.IsNil(latitude) || null.IsNil(longitude) {
		return nil
	}
	nLatitude := strings.TrimSpace(*latitude)
	nLongitude := strings.TrimSpace(*longitude)
	if len(nLatitude) == 0 || len(nLongitude) == 0 {
		return nil
	}
	nLatLong := fmt.Sprintf("%s, %s", nLatitude, nLongitude)
	return &nLatLong
}

func Get(source map[string]interface{}, key *string) (interface{}, bool) {
	values := strings.SplitN(*key, ".", 2)

	if *key == "" {
		return source, true
	} else if len(values) == 1 {
		result, ok := source[*key].(interface{})
		return result, ok
	}

	val, ok := source[values[0]].(map[string]interface{})
	if !ok {
		return nil, false
	}

	return Get(val, &values[1])
}

func ConvertMapStringToStringPointer(payload map[string]string) map[string]*string {
	result := make(map[string]*string)
	for key, value := range payload {
		tempValue := value
		result[key] = &tempValue
	}
	return result
}
