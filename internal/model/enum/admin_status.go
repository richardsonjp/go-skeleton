package enum

import (
	"go-skeleton/pkg/utils/constant"
	"strings"
)

// AdminStatus is connected with Admin model struct #Admin.Status
type AdminStatus int64

//Scan for converting byte to string for fetching/read
func (s *AdminStatus) Scan(value interface{}) error {
	val := value.([]uint8)
	key := string(val)
	for i, v := range adminStatusKey {
		if v == key {
			*s = i
		}
	}
	return nil
}

func NewAdminStatus(value string) AdminStatus {
	for i, v := range adminStatusKey {
		if v == value {
			return i
		}
	}
	panic("enum not found")
}

const (
	ADMIN_ACTIVE AdminStatus = iota + 1
	ADMIN_INACTIVE
)

var adminStatusKey = map[AdminStatus]string{
	ADMIN_ACTIVE:   "active",
	ADMIN_INACTIVE: "inactive",
}

//String for stringify AdminStatus
func (s AdminStatus) String() string {
	return adminStatusKey[s]
}

func GetAdminStatusKeyValuePairs() []constant.KeyValue {
	arr := []constant.KeyValue{}
	for _, v := range adminStatusKey {
		arr = append(arr, constant.KeyValue{
			Key:   v,                // act as key
			Value: strings.Title(v), // act as label
		})
	}
	return arr
}
