package http

import (
	"go-skeleton/pkg/utils/errors"
	"net/http"

	"github.com/stretchr/testify/mock"
)

type HTTPMock struct {
	Mock mock.Mock
}

func (hr *HTTPMock) Get(path string, params *map[string]string, headers map[string]string, cookies []*HTTPCookie, val interface{}) (int, interface{}, []*http.Cookie, error) {
	args := hr.Mock.Called(path, params, headers, cookies)
	if args.Get(0) != nil {
		_ = args.Get(0)
		return 200, nil, nil, nil
	} else {
		return http.StatusInternalServerError, nil, nil, errors.NewGenericError(http.StatusInternalServerError)
	}
}

func (hr *HTTPMock) POSTSend(path string, contentType string, params []byte, headers map[string]string, cookies []*HTTPCookie, val interface{}) (int, interface{}, []*http.Cookie, error) {
	args := hr.Mock.Called(path, contentType, params, headers, cookies, val)
	if args.Get(0) != nil {
		_ = args.Get(0)
		return 200, nil, nil, nil
	} else {
		return http.StatusInternalServerError, nil, nil, errors.NewGenericError(http.StatusInternalServerError)
	}
}

func (hr *HTTPMock) PUTSend(path string, contentType string, params []byte, headers map[string]string, cookies []*HTTPCookie, val interface{}) (int, interface{}, []*http.Cookie, error) {
	args := hr.Mock.Called(path, contentType, params, headers, cookies, val)
	if args.Get(0) != nil {
		_ = args.Get(0)
		return 200, nil, nil, nil
	} else {
		return http.StatusInternalServerError, nil, nil, errors.NewGenericError(http.StatusInternalServerError)
	}
}

func (hr *HTTPMock) POSTFormData(path string, formData map[string]string, headers map[string]string, cookies []*HTTPCookie, val interface{}) (int, interface{}, []*http.Cookie, error) {
	args := hr.Mock.Called(path, formData, headers, cookies, val)
	if args.Get(0) != nil {
		_ = args.Get(0)
		// Return(200, returnInterface, nil, nil).
		return 200, args.Get(1), nil, nil
	} else {
		return http.StatusInternalServerError, nil, nil, errors.NewGenericError(http.StatusInternalServerError)
	}
}

func (hr *HTTPMock) PUTFormData(path string, formData map[string]string, headers map[string]string, cookies []*HTTPCookie, val interface{}) (int, interface{}, []*http.Cookie, error) {
	args := hr.Mock.Called(path, formData, headers, cookies, val)
	if args.Get(0) != nil {
		_ = args.Get(0)
		return 200, nil, nil, nil
	} else {
		return http.StatusInternalServerError, nil, nil, errors.NewGenericError(http.StatusInternalServerError)
	}
}
