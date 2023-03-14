package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"go-skeleton/config"
	"go-skeleton/pkg/utils/api"
	"go-skeleton/pkg/utils/errors"
	"go-skeleton/pkg/utils/logs"
	netutil "go-skeleton/pkg/utils/net"
	"go-skeleton/pkg/utils/parse"
	stringutil "go-skeleton/pkg/utils/strings"
	timeutil "go-skeleton/pkg/utils/time"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)

// Recovery returns a middleware for a given writer that recovers from any panics and writes a 500 if there was one.
func Recovery(mode string) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := timeutil.Now()

		var stash parse.Stashes
		stash.NewRequestBody(c)

		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				brokenPipe := false
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				if logs.Log != nil {
					end := timeutil.Now()
					latency := end.Sub(t)

					var headers map[string][]string = c.Request.Header
					if c.Request.Header.Get("Client-Id") != "" {
						headers["Client-Id"][0] = stringutil.MaskUUIDV4(c.Request.Header.Get("Client-Id"))
					}
					if c.Request.Header.Get("Client-Secret") != "" {
						headers["Client-Secret"][0] = stringutil.MaskUUIDV4(c.Request.Header.Get("Client-Secret"))
					}
					if c.Request.Header.Get("Authorization") != "" {
						headers["Authorization"][0] = stringutil.MaskUUIDV4(c.Request.Header.Get("Authorization"))
					}

					fields := logs.Fields{
						"client_ip":       netutil.GetClientIpAddress(c),
						"client_os":       c.Request.Header.Get("Client-OS"),
						"client_version":  c.Request.Header.Get("Client-Version"),
						"request_id":      c.GetString("RequestId"),
						"request_uri":     c.Request.RequestURI,
						"method":          c.Request.Method,
						"handler":         c.HandlerName(),
						"user_agent":      c.Request.UserAgent(),
						"referer":         c.Request.Referer(),
						"mode":            config.Config.System.Mode,
						"host":            c.Request.Host,
						"path":            c.Request.URL.Path,
						"params":          c.Request.URL.RawQuery,
						"lang":            c.Request.Header.Get("Accept-Language"),
						"status":          http.StatusInternalServerError,
						"process_time":    latency.String(),
						"process_time_ns": latency.Nanoseconds(),
						"error_string":    assertErr(err),
						"error_stack":     errors.GetStack(err),
						"request_body":    stash.GetRequestBody(c),
						"request_header":  c.Request.Header,
						"type_str":        "ERR-PANIC",
					}

					cl := logs.Log.WithFields(fields)
					routePathParamMap := make(map[string]interface{})
					for _, p := range c.Params {
						routePathParamMap[p.Key] = p.Value
					}
					cl = cl.WithFields(logs.Fields{
						"route_path_params": routePathParamMap,
					})

					logs.PushPanicLog(cl)
					c.AbortWithStatusJSON(500, api.Error{
						Message: errors.Translate(c, errors.INTERNAL_SERVER_ERROR),
					})
				}

				// If the connection is dead, we can't write a status to it.
				if brokenPipe {
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
				} else {
					c.AbortWithStatus(http.StatusInternalServerError)
				}
			}
		}()
		c.Next()
	}
}

// stack returns a nicely formatted stack frame, skipping skip frames.
func stack(skip int) []byte {
	buf := &bytes.Buffer{} // the returned data
	// As we loop, we open files and read them. These variables record the currently
	// loaded file.
	var lines [][]byte
	var lastFile string
	for i := skip; ; i++ { // Skip the expected number of frames
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		// Print this much at least.  If we can't find the source, it won't show.
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		if file != lastFile {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
	}
	return buf.Bytes()
}

// source returns a space-trimmed slice of the n'th line.
func source(lines [][]byte, n int) []byte {
	n-- // in stack trace, lines are 1-indexed but our array is 0-indexed
	if n < 0 || n >= len(lines) {
		return dunno
	}
	return bytes.TrimSpace(lines[n])
}

// function returns, if possible, the name of the function containing the PC.
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return dunno
	}
	name := []byte(fn.Name())
	// The name includes the path name to the package, which is unnecessary
	// since the file name is already included.  Plus, it has center dots.
	// That is, we see
	//	runtime/debug.*T·ptrmethod
	// and want
	//	*T.ptrmethod
	// Also the package path might contains dot (e.g. code.google.com/...),
	// so first eliminate the path prefix
	if lastSlash := bytes.LastIndex(name, slash); lastSlash >= 0 {
		name = name[lastSlash+1:]
	}
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, centerDot, dot, -1)
	return name
}

func timeFormat(t time.Time) string {
	timeString := t.Format("2006/01/02 - 15:04:05")
	return timeString
}

func assertErr(err interface{}) string {
	if genericError, ok := err.(*errors.GenericError); ok {
		return genericError.GetErrorDataMessageKey()
	}
	if e, ok := (err).(error); ok {
		return e.Error()
	}
	data, _ := json.Marshal(err)
	return string(data)
}
