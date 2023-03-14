package lang

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle
var CurrentTranslation *Translation

type Translation struct {
	langCode      string
	translateFunc func(key string, templateData interface{}) (string, error)
}

func Init(bd *i18n.Bundle) {
	bundle = bd
	defaultLang := language.AmericanEnglish.String()
	CurrentTranslation = &Translation{
		langCode:      defaultLang,
		translateFunc: GetMappingFunc(defaultLang),
	}
}

func (t *Translation) SetTranslation(translateFunc func(key string, templateData interface{}) (string, error)) {
	t.translateFunc = translateFunc
}

func (t *Translation) Translate(key string, args ...interface{}) (string, error) {
	if len(args) > 0 {
		return t.translateFunc(key, args[0])
	}
	return t.translateFunc(key, nil)
}

func GetMappingFunc(langCode string) func(key string, templateData interface{}) (string, error) {
	localizer := i18n.NewLocalizer(bundle, langCode)
	return func(key string, templateData interface{}) (string, error) {
		if key == "" {
			return "", nil
		}

		return localizer.Localize(&i18n.LocalizeConfig{MessageID: key, TemplateData: templateData})
	}
}

func GetTtranslateFunc(langCode string) func(string) (string, error) {
	localizer := i18n.NewLocalizer(bundle, langCode)
	return func(key string) (string, error) {
		if key == "" {
			return "", nil
		}
		return localizer.Localize(&i18n.LocalizeConfig{MessageID: key})
	}
}
