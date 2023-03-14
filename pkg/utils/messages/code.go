package messages

import "github.com/labstack/gommon/log"

// Code represent message
const (
	NULL int = iota + 20000
	OTP_SENT
	CHANGE_PASSWORD_SUCCESS
	CREDENTIAL_REMOVED
	FORGOT_PASSWORD_SUCCESS
)

// KEYS translate error code to i18n key
var KEYS = map[int]string{
	NULL:                    "NULL",
	OTP_SENT:                "OTP_SENT",
	CHANGE_PASSWORD_SUCCESS: "CHANGE_PASSWORD_SUCCESS",
	CREDENTIAL_REMOVED:      "CREDENTIAL_REMOVED",
	FORGOT_PASSWORD_SUCCESS: "FORGOT_PASSWORD_SUCCESS",
}

func NewMessageCode(value string) int {
	for i, v := range KEYS {
		if v == value {
			return i
		}
	}
	log.Error("message code not found -> ", value)
	return NULL
}
