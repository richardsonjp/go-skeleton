package enum

import (
	"go-skeleton/pkg/utils/constant"
	"strings"
)

// GenericStatus is connected with any model struct which needs the active/inactive status
type GenericStatus int64

// Scan for converting byte to string for fetching/read
func (s *GenericStatus) Scan(value interface{}) error {
	key := value.(string)
	for i, v := range genericStatusKey {
		if v == key {
			*s = i
		}
	}
	return nil
}

func NewGenericStatus(value string) GenericStatus {
	for i, v := range genericStatusKey {
		if v == value {
			return i
		}
	}
	panic("enum not found")
}

const (
	ACTIVE GenericStatus = iota + 1
	INACTIVE
)

var genericStatusKey = map[GenericStatus]string{
	ACTIVE:   "active",
	INACTIVE: "inactive",
}

// String for stringify GenericStatus
func (s GenericStatus) String() string {
	return genericStatusKey[s]
}

func GetGenericStatusKeyValuePairs() []constant.KeyValue {
	var arr []constant.KeyValue
	for _, v := range genericStatusKey {
		arr = append(arr, constant.KeyValue{
			Key:   v,                // act as key
			Value: strings.Title(v), // act as label
		})
	}
	return arr
}
