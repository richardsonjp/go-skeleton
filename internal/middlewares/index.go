package middlewares

import (
	"go-skeleton/pkg/clients/redis"
)

type MiddlewareAccess struct {
	redisDel redis.RedisDelegate
}

func NewMiddlewareAccess(redisDel redis.RedisDelegate) MiddlewareAccess {
	return MiddlewareAccess{
		redisDel: redisDel,
	}
}
