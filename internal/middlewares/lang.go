// middleware / lang.go
package middlewares

import (
	"go-skeleton/pkg/utils/lang"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

func LanguageAccept() gin.HandlerFunc {

	// Validation Accept-Language
	return func(c *gin.Context) {
		langCode := c.GetHeader("Accept-Language")
		switch langCode {
		case "", language.Indonesian.String():
			c.Set("T", lang.GetTtranslateFunc(langCode))
			lang.CurrentTranslation.SetTranslation(lang.GetMappingFunc(langCode))
		default:
			c.Set("T", lang.GetTtranslateFunc(
				language.AmericanEnglish.String(),
			))
			lang.CurrentTranslation.SetTranslation(lang.GetMappingFunc(language.AmericanEnglish.String()))
		}

		c.Next()
	}
}
