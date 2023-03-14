package messages

import (
	goErrors "errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strconv"
	"testing"
)

func mockTranslate(value string) (string, error) {
	return value, nil
}

func mockTranslateFailed(value string) (string, error) {
	return value, goErrors.New("error")
}

func TestTranslateCode(t *testing.T) {
	type Input struct {
		MessageCode  int
		WithContextT bool
		Translate    func(string) (string, error)
	}
	testCases := []struct {
		in   Input
		want string
	}{
		{Input{NULL, true, mockTranslate}, KEYS[NULL]},
		{Input{999999, true, mockTranslate}, fmt.Sprintf("code: %s", strconv.Itoa(999999))},
		{Input{NULL, true, mockTranslateFailed}, fmt.Sprintf("code: %s", strconv.Itoa(NULL))},
		{Input{NULL, true, mockTranslateFailed}, fmt.Sprintf("code: %s", strconv.Itoa(NULL))},
		{Input{NULL, false, mockTranslate}, fmt.Sprintf("code: %s", strconv.Itoa(NULL))},
	}

	for index, tc := range testCases {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		if tc.in.WithContextT {
			c.Set("T", tc.in.Translate)
		}

		result := TranslateCode(c, tc.in.MessageCode)

		assert.Equal(t, tc.want, result, fmt.Sprintf(`TranslateCode: Errir at index %v`, index))
	}
}
