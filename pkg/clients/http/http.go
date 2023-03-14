package http

import (
	"encoding/json"
	"fmt"
	"go-skeleton/pkg/utils/logs"
	stringer "go-skeleton/pkg/utils/strings"
	"net/http"
	"net/url"
	"strings"
	"time"

	timeutil "go-skeleton/pkg/utils/time"

	"github.com/go-resty/resty/v2"
)

const (
	GET    string = "GET"
	POST   string = "POST"
	PUT    string = "PUT"
	PATCH  string = "PATCH"
	UPDATE string = "UPDATE"
	DELETE string = "DELETE"
	JSON   string = "application/json"
	XFORM  string = "application/x-www-form-urlencoded"
	FORM   string = "multipart/form-data"
)

type HTTPDelegate interface {
	Get(path string, params *map[string]string, headers map[string]string, cookies []*HTTPCookie, val interface{}, log bool) (int, interface{}, []*http.Cookie, error)                   // exclusive for GET
	POSTSend(path string, contentType string, params interface{}, headers map[string]string, cookies []*HTTPCookie, val interface{}, log bool) (int, interface{}, []*http.Cookie, error) // POST & PUT (JSON)
	PUTSend(path string, contentType string, params interface{}, headers map[string]string, cookies []*HTTPCookie, val interface{}, log bool) (int, interface{}, []*http.Cookie, error)  // POST & PUT (JSON)
	POSTFormData(path string, formData map[string]string, headers map[string]string, cookies []*HTTPCookie, val interface{}, log bool) (int, interface{}, []*http.Cookie, error)         // POST & PUT (FORM DATA)
	PUTFormData(path string, formData map[string]string, headers map[string]string, cookies []*HTTPCookie, val interface{}, log bool) (int, interface{}, []*http.Cookie, error)          // POST & PUT (FORM DATA)
}

type HTTPClient struct {
	host  string
	conn  *resty.Client
	cache *httpCache
}

type httpCache struct {
	path         string
	params       string
	headers      map[string]string
	responseCode int
	response     interface{}
	contentType  string
	method       string
}

// NewHTTPClient is a constructor function that used to http call.
func NewHTTPClient(host string, duration *time.Duration) HTTPDelegate {
	httpClient := resty.New()
	timeout := time.Duration(5 * time.Second)
	if duration != nil {
		timeout = time.Duration(*duration)
	}
	httpClient.SetTimeout(timeout)

	return &HTTPClient{host: host, conn: httpClient, cache: &httpCache{}}
}

// TraceRequest is a function that used to see last request path, params, headers, content-type, response code and response
func (hr *HTTPClient) TraceRequest() {
	fmt.Printf(
		"Path: %s\nParams: %s\nHeader: %s\nMethod: %s\nContent Type: %s\nResponse Code: %d\nResponse: %s\n",
		hr.cache.path,
		hr.cache.params,
		hr.cache.headers,
		hr.cache.method,
		hr.cache.contentType,
		hr.cache.responseCode,
		hr.cache.response,
	)
}

type HTTPCookie struct {
	Name  string
	Value string
}

func toGenericHTTPCookie(cookie *HTTPCookie) *http.Cookie {
	return &http.Cookie{
		Name:  cookie.Name,
		Value: cookie.Value,
	}
}

func (hr *HTTPClient) getURL(path string) string {

	// Check if path is a URL
	u, err := url.ParseRequestURI(path)
	if err == nil {
		// if error is nill, then parameter path is a url
		return u.String()
	}

	// else, path is valid path
	cleanPath := strings.Trim(path, " ")
	if len(cleanPath) > 0 {
		cleanPath = "/" + cleanPath
	}
	return fmt.Sprintf("%s%s", hr.host, cleanPath)
}

func (hr *HTTPClient) readResponse(response *resty.Response, err error) (interface{}, []*http.Cookie, error) {
	if err != nil {
		return nil, nil, err
	}

	var body interface{}
	if !response.IsError() {
		body = response.Result()
	} else {
		body = response.Error()
	}

	cookie := response.Cookies()

	return body, cookie, err
}

func (hr *HTTPClient) clearCache() {
	hr.cache = &httpCache{}
}

func (hr *HTTPClient) cacheRequest(cache *httpCache, path string, params string, headers *map[string]string, contentType string, method string) {
	cache.path = path
	cache.params = params
	if headers != nil {
		cache.headers = *headers
	}
	cache.contentType = contentType
	cache.method = method
}

func (hr *HTTPClient) cacheResponse(cache *httpCache, responseCode int, response interface{}) {
	cache.responseCode = responseCode
	cache.response = response

	hr.conn.R()
}

func (hr *HTTPClient) get(path string, params *map[string]string, headers map[string]string, cookies []*HTTPCookie, val interface{}) (int, interface{}, []*http.Cookie, error) {
	var paramString string
	hr.clearCache()

	url := hr.getURL(path)

	// setup resty
	req := hr.conn.R().
		SetHeaders(headers).
		SetResult(val)

	// Set Query params
	if params == nil {
		// Handle nil params
		params = &map[string]string{}
	}

	if params != nil {
		req.SetQueryParams(*params)
	}

	// set cookies
	if len(cookies) > 0 {
		var arrCookies []*http.Cookie
		for _, v := range cookies {
			arrCookies = append(arrCookies, toGenericHTTPCookie(v))
		}
		req.SetCookies(arrCookies)
	}

	if params != nil {
		paramString = stringer.ConvertMapToString(*params)
	}

	hr.cacheRequest(hr.cache, path, paramString, &headers, "", GET)

	response, err := req.Execute(GET, url)
	hr.cacheResponse(hr.cache, response.StatusCode(), response.Body())

	if err != nil && response != nil {
		return response.StatusCode(), nil, nil, err
	} else if response == nil {
		return http.StatusInternalServerError, nil, nil, err
	}

	resp, respCookies, err := hr.readResponse(response, err)

	return response.StatusCode(), resp, respCookies, err
}

func (hr *HTTPClient) send(path string, method string, contentType string, params []byte, headers map[string]string, cookies []*HTTPCookie, val interface{}) (int, interface{}, []*http.Cookie, error) {
	hr.clearCache()
	url := hr.getURL(path)

	// setup resty
	req := hr.conn.R().
		SetHeaders(headers).
		SetHeader("Content-Type", contentType).
		SetBody(params)

	if val != nil {
		req.SetResult(val)

		// There is a case that response is not application-json but we need it in JSON format.
		// Impact to our custom struct wont be automatically filled by req.SetResult(val)
		// Use this to automaticaly force resty to transform response content type to application-json
		// and resty will successfully set it
		// req.ForceContentType("application/json")

		req.SetError(val)
	}

	// set cookies
	if len(cookies) > 0 {
		var arrCookies []*http.Cookie
		for _, v := range cookies {
			arrCookies = append(arrCookies, toGenericHTTPCookie(v))
		}
		req.SetCookies(arrCookies)
	}

	hr.cacheRequest(hr.cache, path, string(params), &headers, contentType, method)

	response, err := req.Execute(method, url)
	hr.cacheResponse(hr.cache, response.StatusCode(), response.Body())

	if err != nil && response != nil {
		return response.StatusCode(), nil, nil, err
	} else if response == nil {
		return http.StatusInternalServerError, nil, nil, err
	}

	resp, respCookies, err := hr.readResponse(response, err)

	return response.StatusCode(), resp, respCookies, err
}

func (hr *HTTPClient) sendFormData(path string, method string, formData map[string]string, headers map[string]string, cookies []*HTTPCookie, val interface{}) (int, interface{}, []*http.Cookie, error) {
	hr.clearCache()
	url := hr.getURL(path)

	// setup resty
	req := hr.conn.R().
		SetHeaders(headers).
		SetFormData(formData).
		SetResult(val)

	// set cookies
	if len(cookies) > 0 {
		var arrCookies []*http.Cookie
		for _, v := range cookies {
			arrCookies = append(arrCookies, toGenericHTTPCookie(v))
		}
		req.SetCookies(arrCookies)
	}

	hr.cacheRequest(hr.cache, path, stringer.ConvertMapToString(formData), &headers, XFORM, method)

	response, err := req.Execute(method, url)
	hr.cacheResponse(hr.cache, response.StatusCode(), response.Body())

	if err != nil && response != nil {
		return response.StatusCode(), nil, nil, err
	} else if response == nil {
		return http.StatusInternalServerError, nil, nil, err
	}

	resp, respCookies, err := hr.readResponse(response, err)

	return response.StatusCode(), resp, respCookies, err
}

func (hr *HTTPClient) Get(path string, params *map[string]string, headers map[string]string, cookies []*HTTPCookie, val interface{}, log bool) (int, interface{}, []*http.Cookie, error) {
	startTime := timeutil.Now()
	code, resp, respCookies, err := hr.get(path, params, headers, cookies, val)
	endTime := timeutil.Now()
	if log {
		go httplog(code, hr.getURL(path), GET, params, hr.cache.response, startTime, endTime, err)
	}

	if err != nil {
		return code, nil, nil, err
	}

	return code, resp, respCookies, err
}

func (hr *HTTPClient) POSTSend(path string, contentType string, params interface{}, headers map[string]string, cookies []*HTTPCookie, val interface{}, log bool) (int, interface{}, []*http.Cookie, error) {
	paramData, _ := json.Marshal(params)
	startTime := timeutil.Now()
	code, resp, respCookies, err := hr.send(path, POST, contentType, paramData, headers, cookies, val)
	endTime := timeutil.Now()
	if log {
		go httplog(code, hr.getURL(path), POST, paramData, hr.cache.response, startTime, endTime, err)
	}

	if err != nil {
		return code, nil, nil, err
	}

	return code, resp, respCookies, err
}

func (hr *HTTPClient) PUTSend(path string, contentType string, params interface{}, headers map[string]string, cookies []*HTTPCookie, val interface{}, log bool) (int, interface{}, []*http.Cookie, error) {
	paramData, _ := json.Marshal(params)
	startTime := timeutil.Now()
	code, resp, respCookies, err := hr.send(path, PUT, contentType, paramData, headers, cookies, val)
	endTime := timeutil.Now()
	if log {
		go httplog(code, hr.getURL(path), PUT, paramData, hr.cache.response, startTime, endTime, err)
	}

	if err != nil {
		return code, nil, nil, err
	}

	return code, resp, respCookies, err
}

func (hr *HTTPClient) POSTFormData(path string, formData map[string]string, headers map[string]string, cookies []*HTTPCookie, val interface{}, log bool) (int, interface{}, []*http.Cookie, error) {
	startTime := timeutil.Now()
	code, resp, respCookies, err := hr.sendFormData(path, POST, formData, headers, cookies, val)
	endTime := timeutil.Now()
	if log {
		go httplog(code, hr.getURL(path), POST, formData, hr.cache.response, startTime, endTime, err)
	}

	if err != nil {
		return code, nil, nil, err
	}

	return code, resp, respCookies, err
}

func (hr *HTTPClient) PUTFormData(path string, formData map[string]string, headers map[string]string, cookies []*HTTPCookie, val interface{}, log bool) (int, interface{}, []*http.Cookie, error) {
	startTime := timeutil.Now()
	code, resp, respCookies, err := hr.sendFormData(path, PUT, formData, headers, cookies, val)
	endTime := timeutil.Now()
	if log {
		go httplog(code, hr.getURL(path), PUT, formData, hr.cache.response, startTime, endTime, err)
	}

	if err != nil {
		return code, nil, nil, err
	}

	return code, resp, respCookies, err
}

func httplog(statusCode int, url string, method string, request interface{}, response interface{}, startTime time.Time, endTime time.Time, err error) {
	latency := endTime.Sub(startTime)
	fields := logs.Fields{
		"process_start_time": startTime.String(),
		"process_end_time":   endTime.String(),
		"process_time":       latency.String(),
		"process_time_ns":    latency.Nanoseconds(),
		"status_code":        statusCode,
		"method":             method,
		"route_path":         url,
		"request":            fixConstruct(request),
		"response":           fixConstruct(response),
		"error":              err,
	}

	cl := logs.Log.WithFields(fields)
	cl.Info("HTTP log")
	logs.PushLog("loanhub_http", cl)
}

func fixConstruct(data interface{}) interface{} {
	switch data := data.(type) {
	case []uint8:
		m := make(map[string]interface{})
		_ = json.Unmarshal(data, &m)
		return m
	default:
		return data
	}
}
