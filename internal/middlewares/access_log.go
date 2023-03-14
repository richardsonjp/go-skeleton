package middlewares

import (
	"go-skeleton/config"
	"go-skeleton/pkg/utils/logs"
	netutil "go-skeleton/pkg/utils/net"
	"go-skeleton/pkg/utils/parse"
	stringer "go-skeleton/pkg/utils/strings"
	timeutil "go-skeleton/pkg/utils/time"

	"github.com/gin-gonic/gin"
)

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := timeutil.Now()

		var stash parse.Stashes
		stash.NewRequestBody(c)

		c.Next()

		end := timeutil.Now()
		latency := end.Sub(t)

		var headers map[string][]string = c.Request.Header

		if c.Request.Header.Get("Client-Id") != "" {
			headers["Client-Id"][0] = stringer.MaskUUIDV4(c.Request.Header.Get("Client-Id"))
		}
		if c.Request.Header.Get("Client-Secret") != "" {
			headers["Client-Secret"][0] = stringer.MaskUUIDV4(c.Request.Header.Get("Client-Secret"))
		}
		if c.Request.Header.Get("Authorization") != "" {
			headers["Authorization"][0] = stringer.MaskUUIDV4(c.Request.Header.Get("Authorization"))
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
			"status":          c.Writer.Status(),
			"process_time":    latency.String(),
			"process_time_ns": latency.Nanoseconds(),
			"request_body":    stash.GetRequestBody(c),
			"request_header":  c.Request.Header,
			"type_str":        "GIN",
		}

		cl := logs.Log.WithFields(fields)
		routePathParamMap := make(map[string]interface{})
		for _, p := range c.Params {
			routePathParamMap[p.Key] = p.Value
		}
		cl = cl.WithFields(logs.Fields{
			"route_path_params": routePathParamMap,
		})

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			var errString string
			for i, v := range c.Errors {
				if i >= 1 {
					errString += " | "
				}
				if assertErr(v.Err) != "FDC_NOT_ELIGIBLE" {
					errString += assertErr(v.Err)
				}
			}
			if errString == "" {
				return
			}
			cl = cl.WithFields(logs.Fields{
				"error_string": errString,
			})
			logs.PushLog("loanhub_error", cl)
		} else {
			if c.Request.Method == "GET" || c.Request.Method == "OPTIONS" {
				return
			}

			cl.Info("GIN access log")
			logs.ActivityLog(cl)
		}
	}
}

func GetRoutePath(c *gin.Context) string {

	v, _ := c.Get(logs.SetRoutePath)
	if v == nil {
		return c.Request.URL.Path
	}
	return v.(string)
}
