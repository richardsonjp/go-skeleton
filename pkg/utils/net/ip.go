package net

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// Get client IP Address but without checking IP#IsGlobalUnicast().
// Do not use context#ClientIP it will return invalid IP Address.
// @see https://husobee.github.io/golang/ip-address/2015/12/17/remote-ip-go.html
func GetClientIpAddress(ctx *gin.Context) *string {
	for _, h := range []string{"X-Forwarded-For", "X-Real-Ip"} {
		for _, v := range strings.Split(ctx.Request.Header.Get(h), ",") {
			ip := strings.TrimSpace(v)
			if len(ip) > 0 {
				return &ip
			}
		}
	}
	return GetClientRemoteAddress(ctx)
}

// Get remote adress from context
func GetClientRemoteAddress(ctx *gin.Context) *string {
	remoteAddr := ctx.Request.RemoteAddr
	if len(strings.TrimSpace(remoteAddr)) == 0 {
		return nil
	}
	remoteAddr = strings.TrimSpace(remoteAddr)
	return &remoteAddr
}
