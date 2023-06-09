package null

import (
	"reflect"
)

func NullAsStringEmpty(value *string) string {
	if value == nil {
		return ""
	}

	return *value
}

func StringEmptyAsNull(value string) *string {
	if value == "" {
		return nil
	}

	return &value
}

func IsNil(value interface{}) bool {
	if value == nil {
		return true
	}
	if reflect.TypeOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil() {
		return true
	}
	return false
}
