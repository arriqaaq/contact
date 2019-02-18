package api

import (
	"fmt"
	"github.com/arriqaaq/zizou"
	"github.com/go-redis/redis"
	"time"
)

type Cache interface {
	Get(string) (interface{}, bool)
	Set(string, interface{}, time.Duration) error
}

func NewL1Cache() Cache {
	zizouCnf := &zizou.Config{
		SweepTime: DefaultInternalCacheEvictionTime,
		ShardSize: 256,
	}
	l1cache, _ := zizou.New(zizouCnf)
	return l1cache
}

func NewRdbConfig() RdbConfig {
	return RdbConfig{
		EnableLog: true,
	}
}

// Redis configuration
type RdbConfig struct {
	EnableLog bool
	redis.Options
}

type redisDb struct {
	conn *redis.Client
	opts *RdbConfig
}

func NewRedisCache(opt *RdbConfig) (Cache, error) {
	n := &redisDb{opts: opt}
	opts := opt.Options
	client := redis.NewClient(&opts)
	n.conn = client
	_, err := client.Ping().Result()
	return n, err
}

func (r *redisDb) Get(key string) (interface{}, bool) {
	val, err := r.conn.Get(key).Result()
	if err != nil {
		if r.opts.EnableLog {
			fmt.Println("redis:get:error: ", err)
		}
		return val, false
	}
	return val, true
}

func (r *redisDb) Set(key string, value interface{}, duration time.Duration) error {
	return r.conn.Set(key, value, duration).Err()
}
