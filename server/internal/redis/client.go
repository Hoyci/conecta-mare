package redis

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewClient(host, port string) *redis.Client {
	addr := fmt.Sprintf("%s:%s", host, port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	return rdb
}
