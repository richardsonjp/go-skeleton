package redis

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"go-skeleton/config"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

type RedisDelegate interface {
	Init()
	InitNoUse()
	Use() *redis.Client
	Get(key string, val interface{}) error
	Set(key string, val interface{}, duration time.Duration) error
	GetFirstMatchedKey(keyPattern string) (string, error)
	GetFirstMatchedValue(keyPattern string, val interface{}) error
	Delete(key string) error
}

type redisDelegate struct {
	redisConn *redis.Client
	once      sync.Once
}

func NewRedisDel() RedisDelegate {
	return &redisDelegate{}
}

// Init get client
func (redisDel *redisDelegate) Init() {
	redisDel.run(true)
}

// InitNoUse no client
func (redisDel *redisDelegate) InitNoUse() {
	redisDel.run(false)
}

// Get() parameter 'val' must be cast with `var example = &struct{}` , (value and struct must be cast exactly)
func (redisDel *redisDelegate) Get(key string, val interface{}) error {
	p, err := redisDel.redisConn.Get(key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(p, val)
}

func (redisDel *redisDelegate) Delete(key string) error {
	v := redisDel.redisConn.Get(key)
	if v.Err() == redis.Nil {
		return nil
	}
	_ = redisDel.redisConn.Del(key)
	return nil
}

/**
 * @return First matched key
 */
func (redisDel *redisDelegate) GetFirstMatchedKey(keyPattern string) (string, error) {
	var cursor uint64
	var targetKey string
	iter := redisDel.redisConn.Scan(cursor, keyPattern, 0).Iterator()
	for iter.Next() {
		targetKey = iter.Val()
		break
	}
	if err := iter.Err(); err != nil {
		return targetKey, err
	}
	return targetKey, nil
}

/**
 * @return First value matched by key pattern
 */
func (redisDel *redisDelegate) GetFirstMatchedValue(keyPattern string, val interface{}) error {
	matchedKey, err := redisDel.GetFirstMatchedKey(keyPattern)
	if err != nil {
		return err
	}
	return redisDel.Get(matchedKey, val)
}

func (redisDel *redisDelegate) Set(key string, val interface{}, duration time.Duration) error {
	v, _ := json.Marshal(val)
	err := redisDel.redisConn.Set(key, v, duration).Err()
	if err != nil {
		return err
	}
	return nil
}

// Use Redis connection
func (redisDel *redisDelegate) Use() *redis.Client {
	return redisDel.redisConn
}

func (redisDel *redisDelegate) run(use bool) {
	if use {
		config := config.Config.Redis
		redisDel.once.Do(func() {
			redisOptions := &redis.Options{
				Addr:     config.Address,
				Password: config.Password,
				DB:       config.DB, //default
			}
			if config.TLSEnabled {
				redisOptions.TLSConfig = &tls.Config{
					InsecureSkipVerify: true,
				}
			}

			redisDel.redisConn = redis.NewClient(redisOptions)

			_, err := redisDel.redisConn.Ping().Result()
			if err != nil {
				panic("init redis failed: " + err.Error())
			}

		})
	} else {
		fmt.Println("redis not used")
	}

}
