package errors

import (
	"go-skeleton/pkg/utils/constant"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

// Code represent error
const (
	NULL int = iota + 1000
	DATA_NOT_FOUND
	CLIENT_AUTH_ERROR
	CLIENT_AUTH_FORBIDDEN
	CLIENT_INFO_NOT_FOUND
	INVALID_ENCRYPTION
	UNPROCESSABLE_ERROR
	FAILED_BINDING_PAYLOAD
	INTERNAL_SERVER_ERROR
	UNPROCESSABLE_REQUEST
	SERVICE_DEPENDENCY_ERROR
	DUPLICATE_DATA
	DUPLICATE_REGISTRATION_DATA
	VALIDATION_ERROR
	AUTHENTICATION_FAIL
	TOKEN_UNAUTHORIZED
	INVALID_EMAIL_OR_PASSWORD
	INVALID_OLD_PASSWORD
	USER_REMOVED
	INVALID_RECAPTCHA
	INVALID_RECAPTCHA_DUPLICATE
	UNAUTHORIZED_REQUEST
	PARAM_NOT_VIABLE
	CLIENT_ERROR
	INVALID_IP_ADDRESS
	FORBIDDEN
	FORGOT_PASSWORD_EXCEED_LIMIT
	CALLBACK_ALREADY_SENT
	EXCEEDED_MAX_DATE_RANGE
	BAD_GATEWAY
)

// ERROR_KEYS translate error code to i18n key, determine http status code and error code shown to client
var ERROR_KEYS = map[int]ErrorData{
	NULL:                         *NewErrorData("NULL", http.StatusBadRequest, ERROR_VALIDATION),
	DATA_NOT_FOUND:               *NewErrorData("DATA_NOT_FOUND", http.StatusBadRequest, ERROR_VALIDATION),
	CLIENT_AUTH_ERROR:            *NewErrorData("CLIENT_AUTH_ERROR", http.StatusUnauthorized, ERROR_VALIDATION),
	CLIENT_AUTH_FORBIDDEN:        *NewErrorData("CLIENT_AUTH_FORBIDDEN", http.StatusForbidden, ERROR_VALIDATION),
	CLIENT_INFO_NOT_FOUND:        *NewErrorData("CLIENT_INFO_NOT_FOUND", http.StatusBadRequest, ERROR_VALIDATION),
	INVALID_ENCRYPTION:           *NewErrorData("INVALID_ENCRYPTION", http.StatusBadRequest, ERROR_VALIDATION),
	UNPROCESSABLE_ERROR:          *NewErrorData("UNPROCESSABLE_ERROR", http.StatusInternalServerError, ERROR_UNKNOWN),
	FAILED_BINDING_PAYLOAD:       *NewErrorData("FAILED_BINDING_PAYLOAD", http.StatusBadRequest, ERROR_VALIDATION),
	INTERNAL_SERVER_ERROR:        *NewErrorData("INTERNAL_SERVER_ERROR", http.StatusBadRequest, ERROR_VALIDATION),
	UNPROCESSABLE_REQUEST:        *NewErrorData("UNPROCESSABLE_REQUEST", http.StatusBadRequest, ERROR_VALIDATION),
	SERVICE_DEPENDENCY_ERROR:     *NewErrorData("SERVICE_DEPENDENCY_ERROR", http.StatusBadRequest, ERROR_VALIDATION),
	VALIDATION_ERROR:             *NewErrorData("VALIDATION_ERROR", http.StatusBadRequest, ERROR_VALIDATION),
	DUPLICATE_DATA:               *NewErrorData("DUPLICATE_DATA", http.StatusBadRequest, ERROR_VALIDATION),
	DUPLICATE_REGISTRATION_DATA:  *NewErrorData("DUPLICATE_REGISTRATION_DATA", http.StatusBadRequest, ERROR_VALIDATION),
	AUTHENTICATION_FAIL:          *NewErrorData("AUTHENTICATION_FAIL", http.StatusUnauthorized, ERROR_NULL),
	TOKEN_UNAUTHORIZED:           *NewErrorData("TOKEN_UNAUTHORIZED", http.StatusBadRequest, ERROR_VALIDATION),
	USER_REMOVED:                 *NewErrorData("USER_REMOVED", http.StatusBadRequest, ERROR_VALIDATION),
	INVALID_EMAIL_OR_PASSWORD:    *NewErrorData("INVALID_EMAIL_OR_PASSWORD", http.StatusBadRequest, ERROR_VALIDATION),
	INVALID_OLD_PASSWORD:         *NewErrorData("INVALID_OLD_PASSWORD", http.StatusBadRequest, ERROR_VALIDATION),
	INVALID_RECAPTCHA:            *NewErrorData("INVALID_RECAPTCHA", http.StatusBadRequest, ERROR_VALIDATION),
	INVALID_RECAPTCHA_DUPLICATE:  *NewErrorData("INVALID_RECAPTCHA_DUPLICATE", http.StatusBadRequest, ERROR_VALIDATION),
	UNAUTHORIZED_REQUEST:         *NewErrorData("UNAUTHORIZED_REQUEST", http.StatusUnauthorized, ERROR_VALIDATION),
	PARAM_NOT_VIABLE:             *NewErrorData("PARAM_NOT_VIABLE", http.StatusBadRequest, ERROR_VALIDATION),
	CLIENT_ERROR:                 *NewErrorData("CLIENT_ERROR", http.StatusBadRequest, ERROR_VALIDATION),
	INVALID_IP_ADDRESS:           *NewErrorData("INVALID_IP_ADDRESS", http.StatusBadRequest, ERROR_VALIDATION),
	FORBIDDEN:                    *NewErrorData("FORBIDDEN", http.StatusForbidden, ERROR_NULL),
	FORGOT_PASSWORD_EXCEED_LIMIT: *NewErrorData("FORGOT_PASSWORD_EXCEED_LIMIT", http.StatusBadRequest, ERROR_VALIDATION),
	CALLBACK_ALREADY_SENT:        *NewErrorData("CALLBACK_ALREADY_SENT", http.StatusBadRequest, ERROR_VALIDATION),
	EXCEEDED_MAX_DATE_RANGE:      *NewErrorData("EXCEEDED_MAX_DATE_RANGE", http.StatusBadRequest, ERROR_VALIDATION),
	BAD_GATEWAY:                  *NewErrorData("BAD_GATEWAY", http.StatusBadRequest, ERROR_VALIDATION),
}

type ErrorData struct {
	// MessageKey is the key that will be used to translate the message. mapping key to id/en.json
	MessageKey string
	// HttpCode is the HTTP status code to be returned to the client.
	HttpCode int
	// ErrorCode is the error code that will be used to identify specific error.
	ErrorCode constant.ReserveErrorCode
}

func NewErrorData(messageKey string, httpCode int, errorCode constant.ReserveErrorCode) *ErrorData {
	return &ErrorData{
		MessageKey: messageKey,
		HttpCode:   httpCode,
		ErrorCode:  errorCode,
	}
}

// ERROR_ is reserved key for error code
const (
	ERROR_NULL    constant.ReserveErrorCode = 0
	ERROR_UNKNOWN constant.ReserveErrorCode = 1
)

// ERROR_ is reserved key for error code
const (
	ERROR_VALIDATION constant.ReserveErrorCode = iota + 4000
)

var (
	ResourceNil error = errors.New("create resource nil")
)

type Err string

func CustomError(message string) error {
	return errors.New(message)
}

// Loanhub error wrapper.
// @TODO Handling thread-safe translation without providing context
type GenericError struct {
	// code that refers to enum iota of "Code represent error"
	code int

	// override default error message
	// message will "ignore/override" translated message from code
	message string

	// errors is the additional error data that shown to client
	details []constant.ErrorDetails

	// callback after the message is translated
	fn func(string) string
}

func (se *GenericError) Error() string {
	return strconv.Itoa(se.GetCode())
}

func (se *GenericError) GetCode() int {
	return se.code
}

func (se *GenericError) GetMessage() string {
	return se.message
}

func (se *GenericError) GetCallback() func(string) string {
	return se.fn
}

func (se *GenericError) GetDetails() []constant.ErrorDetails {
	if se.details == nil {
		return []constant.ErrorDetails{}
	}
	return se.details
}

func (se *GenericError) GetErrorDataCode() constant.ReserveErrorCode {
	return ERROR_KEYS[se.code].ErrorCode
}

func (se *GenericError) GetErrorDataMessageKey() string {
	return ERROR_KEYS[se.code].MessageKey
}

func NewGenericError(code int, options ...func(*GenericError)) error {
	genericError := &GenericError{code: code}

	for _, option := range options {
		option(genericError)
	}

	return genericError
}

func WithCallback(fn func(string) string) func(err *GenericError) {
	return func(err *GenericError) {
		err.fn = fn
	}
}

func OverrideCode(code int) func(err *GenericError) {
	return func(err *GenericError) {
		err.code = code
	}
}

func OverrideErrorMessage(message string) func(err *GenericError) {
	return func(err *GenericError) {
		err.message = message
	}
}

func SetDetails(details []constant.ErrorDetails) func(err *GenericError) {
	return func(err *GenericError) {
		err.details = details
	}
}

// NewGenericErrorWithFn creates a custom error
// Deprecated: use NewGenericError with "OverrideErrorMessage" options instead
// will be removed in the future
func NewGenericErrorWithFn(code int, fn func(string) string) error {
	return &GenericError{code: code, fn: fn}
}
