package rdx

import (
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
)

func Open(c *Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Addr,
		Password:     c.Password,
		DB:           int(c.Db),
		DialTimeout:  c.DialTimeout.AsDuration(),
		WriteTimeout: c.WriteTimeout.AsDuration(),
		ReadTimeout:  c.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	return rdb
}
