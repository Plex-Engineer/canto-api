package redisclient

import (
	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func NewClient() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,
	})
}
