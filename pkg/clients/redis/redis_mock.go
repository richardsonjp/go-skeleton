package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/mock"
)

type RedisMock struct {
	Mock mock.Mock
}

func (o *RedisMock) Init() {
	o.run(true)
}

func (o *RedisMock) InitNoUse() {
	o.run(false)
}

func (o *RedisMock) Get(key string, val interface{}) error {
	args := o.Mock.Called(key, val)
	if args.Get(0) != nil {
		val = args.Get(0)
		return nil
	} else {
		return args.Error(0)
	}
}

func (o *RedisMock) GetFirstMatchedKey(keyPattern string) (string, error) {
	args := o.Mock.Called(keyPattern)
	if args.Get(0) != nil {
		str := fmt.Sprintf("%v", args.Get(0))
		return str, nil
	} else {
		return "", args.Error(1)
	}
}

func (o *RedisMock) GetFirstMatchedValue(keyPattern string, val interface{}) error {
	args := o.Mock.Called(keyPattern, val)
	if args.Get(0) != nil {
		return nil
	} else {
		return args.Error(1)
	}
}

func (o *RedisMock) Set(key string, val interface{}, duration time.Duration) error {
	args := o.Mock.Called(key, val, duration)
	if args.Error(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (o *RedisMock) Use() *redis.Client {
	panic("@TODO not implemented")
}

func (o *RedisMock) run(use bool) {
	panic("@TODO not implemented")
}
