package redis

import (
	"github.com/redis/go-redis/v9"
)

type Client struct {
	db *redis.Client
}

func GetClient() Client {
	return Client{rdb}
}
