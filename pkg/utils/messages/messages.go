package messages

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TranslateCode(ctx *gin.Context, messageCode int) string {
	T, ok := ctx.Get("T")
	if ok {
		translator, ok := T.(func(string) (string, error))
		if ok {
			translatedMessage, err := translator(KEYS[messageCode])
			if err == nil && len(translatedMessage) > 0 {
				return translatedMessage
			} else {
				return fmt.Sprintf("code: %s", strconv.Itoa(messageCode))
			}
		} else {
			return fmt.Sprintf("code: %s", strconv.Itoa(messageCode))
		}
	} else {
		return fmt.Sprintf("code: %s", strconv.Itoa(messageCode))
	}
}
