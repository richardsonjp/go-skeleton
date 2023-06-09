package validator

import (
	"github.com/go-playground/validator/v10"
)

type MessageOverwriter struct {
	// Validator key
	Key string
	// field: value
	JsonCounterPart map[string]string
	// Return validation message
	GetMessage func(e validator.FieldError, mo MessageOverwriter) string
}

var MessageOverwriters = []MessageOverwriter{
	*neField,
	*eqField,
}

func findMessageOverwriter(tag string, field string) *MessageOverwriter {
	for _, v := range MessageOverwriters {
		if v.Key == tag {
			if _, ok := v.JsonCounterPart[field]; ok {
				return &v
			}
		}
	}
	return nil
}

var neField = &MessageOverwriter{
	Key: "nefield",
	JsonCounterPart: map[string]string{
		"new_password": "old_password",
	},
	GetMessage: func(e validator.FieldError, mo MessageOverwriter) string {
		failedField := e.Field()
		counterField := mo.JsonCounterPart[failedField]
		return failedField + " harus berbeda dengan " + counterField
	},
}

var eqField = &MessageOverwriter{
	Key: "eqfield",
	JsonCounterPart: map[string]string{
		"confirm_password": "new_password",
	},
	GetMessage: func(e validator.FieldError, mo MessageOverwriter) string {
		failedField := e.Field()
		counterField := mo.JsonCounterPart[failedField]
		return failedField + " harus sama dengan " + counterField
	},
}
