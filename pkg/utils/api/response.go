package api

import "go-skeleton/pkg/utils/constant"

type (
	Base struct {
		Data interface{} `json:"data"`
	}

	BaseWithMeta struct {
		Data interface{} `json:"data"`
		Meta interface{} `json:"meta"`
	}

	Error struct {
		Code    constant.ReserveErrorCode `json:"code"`
		Message string                    `json:"message"`
		Details []constant.ErrorDetails   `json:"details"`
	}

	Message struct {
		Message string `json:"message"`
	}

	List struct {
		Data interface{} `json:"data"`
	}
)
