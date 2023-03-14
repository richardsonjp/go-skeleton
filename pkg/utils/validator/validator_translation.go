package validator

import (
	"go-skeleton/pkg/utils/lang"

	"github.com/go-playground/validator/v10"
)

type ValidatorTranslation struct {
	// Validator key
	Key string
	// Return validation message
	GetMessage func(e validator.FieldError, mo ValidatorTranslation) string
}

var ValidatorTranslations = []ValidatorTranslation{
	*required,
	*length,
	*email,
	*max,
	*min,
	*numeric,
	*url,
}

func findValidatorTranslation(tag string, field string, param string) *ValidatorTranslation {
	for _, v := range ValidatorTranslations {
		if v.Key == tag {
			return &v
		}
	}
	return nil
}

var required = &ValidatorTranslation{
	Key: "required",
	GetMessage: func(e validator.FieldError, mo ValidatorTranslation) string {
		failedField := e.Field()
		templateData := map[string]interface{}{"Field": failedField}
		msg, _ := lang.CurrentTranslation.Translate("ERR_REQUIRED", templateData)
		return msg
	},
}

var length = &ValidatorTranslation{
	Key: "length",
	GetMessage: func(e validator.FieldError, mo ValidatorTranslation) string {
		failedField := e.Field()
		params := e.Param()
		templateData := map[string]interface{}{"Field": failedField, "Length": params}
		msg, _ := lang.CurrentTranslation.Translate("ERR_LEN", templateData)
		return msg
	},
}

var email = &ValidatorTranslation{
	Key: "email",
	GetMessage: func(e validator.FieldError, mo ValidatorTranslation) string {
		failedField := e.Field()
		templateData := map[string]interface{}{"Field": failedField}
		msg, _ := lang.CurrentTranslation.Translate("ERR_EMAIL", templateData)
		return msg
	},
}

var max = &ValidatorTranslation{
	Key: "max",
	GetMessage: func(e validator.FieldError, mo ValidatorTranslation) string {
		failedField := e.Field()
		params := e.Param()
		templateData := map[string]interface{}{"Field": failedField, "Length": params}
		msg, _ := lang.CurrentTranslation.Translate("ERR_MAX", templateData)
		return msg
	},
}

var min = &ValidatorTranslation{
	Key: "min",
	GetMessage: func(e validator.FieldError, mo ValidatorTranslation) string {
		failedField := e.Field()
		params := e.Param()
		templateData := map[string]interface{}{"Field": failedField, "Length": params}
		msg, _ := lang.CurrentTranslation.Translate("ERR_MIN", templateData)
		return msg
	},
}

var numeric = &ValidatorTranslation{
	Key: "numeric",
	GetMessage: func(e validator.FieldError, mo ValidatorTranslation) string {
		failedField := e.Field()
		templateData := map[string]interface{}{"Field": failedField}
		msg, _ := lang.CurrentTranslation.Translate("ERR_NUMERIC", templateData)
		return msg
	},
}

var url = &ValidatorTranslation{
	Key: "url",
	GetMessage: func(e validator.FieldError, mo ValidatorTranslation) string {
		failedField := e.Field()
		templateData := map[string]interface{}{"Field": failedField}
		msg, _ := lang.CurrentTranslation.Translate("ERR_URL", templateData)
		return msg
	},
}
